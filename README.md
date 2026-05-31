# Alumni Tracert

Sistem pelacakan alumni dan legalisir ijazah digital Poltekkes Medan.

## Tech Stack

| Component | Technology |
|-----------|-----------|
| Backend | Go (gRPC + gRPC-Web) |
| Frontend | Svelte + TailwindCSS |
| Database | MySQL 8.0 |
| Mobile | React Native |
| Container | Docker + Docker Compose |
| SSL | Let's Encrypt (certbot) |
| Domain | https://tracert.ttmi.my.id |

## Arsitektur

```
Browser → https://tracert.ttmi.my.id
  └─ Host nginx (SSL termination)
       └─ Docker frontend:8080 (nginx + Svelte static)
            ├── Static files (HTML/CSS/JS)
            └── /proto.TracertService/* → Docker backend:8099 (gRPC-web)
                 └── MySQL:3306

Mobile → backend:9099 (native gRPC)
```

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Domain pointing to server (for SSL)

### Setup

```bash
# 1. Clone repository
git clone <repo-url> && cd alumni-tracert

# 2. Configure environment
cp .env.example .env
# Edit .env if needed (default values work for local dev)

# 3. Build and start
make docker-up

# 4. Verify
docker compose ps
curl -s https://tracert.ttmi.my.id/
```

### Manual SSL Setup (first time)

```bash
# Nginx config already created at /etc/nginx/sites-available/tracert.ttmi.my.id
# SSL certificate already obtained via certbot

# To renew manually:
sudo certbot renew

# Certbot auto-renews via systemd timer
```

## Akun Default

| Role | Email | Password | user_type |
|------|-------|----------|-----------|
| **Admin** | admin@poltekkes-medan.ac.id | admin123 | 3 |
| **Pejabat** | pejabat@poltekkes-medan.ac.id | pejabat123 | 4 |
| **Alumni** | alumni@poltekkes-medan.ac.id | alumni123 | 1 |
| **Appraiser** | appraiser@poltekkes-medan.ac.id | appraiser123 | 2 |

> ⚠️ Change all passwords before deploying to production!

## Role & Permission

### Alumni (user_type=1)
- Registrasi data alumni (nama, NIK, tempat/tanggal lahir, no. telp)
- Input sertifikat (NIM, no. ijazah, jurusan, tahun masuk/lulus)
- Upload ijazah & transkrip untuk legalisir
- Melihat status legalisir
- Mengisi kuisioner tracer study
- Memberikan rating setelah legalisir selesai

### Appraiser (user_type=2)
- Mengisi penilaian terhadap alumni (posisi, instansi)
- Data penilai disimpan untuk validasi

### Admin (user_type=3)
- Verifikasi pengajuan legalisir
- Manage data alumni dan users
- Export data kuisioner
- Melihat daftar admin dan pejabat
- Reject legalisir dengan alasan

### Pejabat (user_type=4)
- Approve legalisir yang sudah diverifikasi admin
- Tanda tangan digital (upload file signed)
- Hanya bisa approve yang status=verified dan is_offline=FALSE

## Alur Legalisir

```
┌─────────┐    ┌──────────┐    ┌──────────┐    ┌─────────┐
│  Alumni  │───▶│  Submit  │───▶│  Admin   │───▶│ Pejabat │
│  Upload  │    │ status=1 │    │ Verify   │    │ Approve │
└─────────┘    └──────────┘    │ status=2 │    │ status=3│
                               └──────────┘    └─────────┘
                                    │
                                    ▼ (jika ditolak)
                               ┌──────────┐
                               │ Rejected │
                               │ status=0 │
                               └──────────┘
```

### Status Code
| Status | Keterangan |
|--------|-----------|
| 0 | Ditolak (rejected) |
| 1 | Diajukan (submitted) |
| 2 | Diverifikasi (verified) |
| 3 | Disetujui (approved) |

### Jenis Legalisir
- **Online** (`is_offline=FALSE`): Alumni upload → Admin verify → Pejabat approve
- **Offline** (`is_offline=TRUE`): Alumni datang langsung → Admin verify & approve

## Database Schema

### Tables

```
users
├── alumni (1:1 with users where user_type=1)
│   ├── certificates (1:N)
│   │   └── legalizes (1:1)
│   └── alumni_appraisers (N:M with users where user_type=2)
├── tracers (1:N with users)
│   └── user_answers (1:N)
├── question_groups
│   └── questions (1:N)
│       └── question_options (1:N)
└── request_passwords (password reset tokens)
```

### ERD

```
┌──────────┐     ┌──────────┐     ┌──────────────┐
│  users   │────▶│  alumni  │────▶│ certificates │
│          │     │          │     │              │
│ id       │     │ id       │     │ id           │
│ email    │     │ user_id  │     │ alumni_id    │
│ password │     │ name     │     │ nim          │
│ name     │     │ nik      │     │ no_ijazah    │
│ user_type│     │ phone    │     │ major_study  │
└──────────┘     └──────────┘     └──────────────┘
     │                                 │
     │           ┌───────────┐         │
     │           │ legalizes │◀────────┘
     │           │           │
     │           │ id (UUID) │
     │           │ cert_id   │
     │           │ status    │
     │           │ is_offline│
     └──────────▶│ verified  │
                 │ approved  │
                 └───────────┘
```

## API (gRPC)

### Service: TracertService

#### Auth
| Method | Request | Response | Description |
|--------|---------|----------|-------------|
| Login | LoginInput (email, password) | User | Autentikasi user |
| GetUserLogin | Token | User | Ambil data user dari token |
| ChangePassword | ChangePassword | GenericResponse | Ganti password |
| ForgotPassword | RequestPassword | GenericResponse | Request reset password |
| ResetPassword | RequestPassword | GenericResponse | Reset password dengan token |

#### Alumni
| Method | Description |
|--------|-------------|
| AlumniRegistration | Registrasi alumni baru |
| AlumniUpdate | Update data alumni |
| AlumniList | List semua alumni |
| AlumniGet | Detail alumni |

#### Certificate
| Method | Description |
|--------|-------------|
| CertificateInsert | Tambah sertifikat |
| CertificateUpdate | Update sertifikat |
| CertificateList | List sertifikat alumni |

#### Legalize
| Method | Description | Role |
|--------|-------------|------|
| LegalizeInsert | Ajukan legalisir | Alumni |
| LegalizeList | List legalisir | All |
| LegalizeGet | Detail legalisir | All |
| LegalizeVerify | Verifikasi legalisir | Admin |
| LegalizeReject | Tolak legalisir | Admin |
| LegalizeApprove | Approve + tanda tangan | Pejabat |
| LegalizeRate | Beri rating | Alumni |
| LegalizeExtended | Perpanjang masa berlaku | Alumni |

#### Questionnaire
| Method | Description | Role |
|--------|-------------|-------------|
| QuestionGroupList | List pertanyaan kuisioner | All |
| UserAnswerSubmit | Jawab kuisioner | Alumni |
| UserAnswerList | Export jawaban kuisioner | Admin |

## Docker Commands

```bash
# Start all services
make docker-up

# Stop all services
make docker-down

# View logs
make docker-logs

# Stop and remove volumes (destroys DB)
make docker-clean

# Rebuild single service
docker compose up -d --build backend
docker compose up -d --build frontend

# Access MySQL
docker exec -it alumni-mysql mysql -u alumni -palumni_pass alumni_db

# View backend logs
docker compose logs backend -f

# View frontend logs
docker compose logs frontend -f
```

## Port Mapping

| Service | Container Port | Host Binding | Public |
|---------|---------------|--------------|--------|
| MySQL | 3306 | 127.0.0.1:3306 | No |
| Backend gRPC | 9099 | 127.0.0.1:9099 | No |
| Backend gRPC-Web | 8099 | 127.0.0.1:8099 | No |
| Frontend (nginx) | 8080 | 127.0.0.1:8080 | No |
| Host nginx | 80, 443 | 0.0.0.0 | Yes |

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| MYSQL_ROOT_PASSWORD | rootpassword | MySQL root password |
| MYSQL_DATABASE | alumni_db | Database name |
| MYSQL_USER | alumni | Database user |
| MYSQL_PASSWORD | alumni_pass | Database password |
| APP_ENV | production | Application environment |
| DB_DRIVER | mysql | Database driver |
| DB_SOURCE | (auto) | Database connection string |

## Troubleshooting

### Backend tidak bisa connect ke MySQL
```bash
# Cek MySQL healthy
docker compose ps
# Cek logs backend
docker compose logs backend
```

### Frontend blank / JS error
```bash
# Rebuild frontend
docker compose up -d --build frontend
```

### SSL certificate expired
```bash
sudo certbot renew
sudo systemctl reload nginx
```

### Reset database
```bash
docker compose down
docker volume rm alumni-tracert_mysql_data
docker compose up -d
```

## License

[GNU GPL](LICENSE)
