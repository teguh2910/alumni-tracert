-- Alumni Tracert Database Schema
-- Generated from backend Go source code

SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

-- ============================================
-- Users table
-- user_type: 1=alumni, 2=appraiser, 3=admin, 4=pejabat
-- ============================================
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL DEFAULT '',
  `is_actived` TINYINT(1) NOT NULL DEFAULT 1,
  `user_type` INT UNSIGNED NOT NULL DEFAULT 1,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Alumni table
-- ============================================
CREATE TABLE IF NOT EXISTS `alumni` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL DEFAULT '',
  `nik` VARCHAR(20) NOT NULL DEFAULT '',
  `place_of_birth` VARCHAR(255) NOT NULL DEFAULT '',
  `date_of_birth` VARCHAR(50) NOT NULL DEFAULT '',
  `phone` VARCHAR(20) NOT NULL DEFAULT '',
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_alumni_user_id` (`user_id`),
  CONSTRAINT `fk_alumni_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Alumni appraisers table
-- ============================================
CREATE TABLE IF NOT EXISTS `alumni_appraisers` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `alumni_id` BIGINT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL DEFAULT '',
  `instansi` VARCHAR(255) NOT NULL DEFAULT '',
  `position` VARCHAR(255) NOT NULL DEFAULT '',
  `alumni_position` VARCHAR(255) NOT NULL DEFAULT '',
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_appraiser_user_id` (`user_id`),
  KEY `idx_appraiser_alumni_id` (`alumni_id`),
  CONSTRAINT `fk_appraiser_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_appraiser_alumni` FOREIGN KEY (`alumni_id`) REFERENCES `alumni` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Certificates table
-- ============================================
CREATE TABLE IF NOT EXISTS `certificates` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `alumni_id` BIGINT UNSIGNED NOT NULL,
  `nim` VARCHAR(50) NOT NULL DEFAULT '',
  `no_ijazah` VARCHAR(100) NOT NULL DEFAULT '',
  `major_study` VARCHAR(255) NOT NULL DEFAULT '',
  `entry_year` VARCHAR(10) NOT NULL DEFAULT '',
  `graduation_year` VARCHAR(10) NOT NULL DEFAULT '',
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_certificates_alumni_id` (`alumni_id`),
  CONSTRAINT `fk_certificates_alumni` FOREIGN KEY (`alumni_id`) REFERENCES `alumni` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Legalizes table
-- status: 0=rejected, 1=submitted, 2=verified, 3=approved
-- ============================================
CREATE TABLE IF NOT EXISTS `legalizes` (
  `id` VARCHAR(36) NOT NULL,
  `certificate_id` BIGINT UNSIGNED NOT NULL,
  `ijazah` VARCHAR(500) NOT NULL DEFAULT '',
  `transcript` VARCHAR(500) NOT NULL DEFAULT '',
  `is_offline` TINYINT(1) NOT NULL DEFAULT 0,
  `is_verified` TINYINT(1) NOT NULL DEFAULT 0,
  `is_approved` TINYINT(1) NOT NULL DEFAULT 0,
  `is_extend` TINYINT(1) NOT NULL DEFAULT 0,
  `verified_by` BIGINT UNSIGNED DEFAULT NULL,
  `verified_at` DATETIME DEFAULT NULL,
  `approved_by` BIGINT UNSIGNED DEFAULT NULL,
  `approved_at` DATETIME DEFAULT NULL,
  `ijazah_signed` VARCHAR(500) DEFAULT NULL,
  `transcript_signed` VARCHAR(500) DEFAULT NULL,
  `status` INT UNSIGNED NOT NULL DEFAULT 1,
  `rating` INT UNSIGNED DEFAULT NULL,
  `rejected_reason` TEXT DEFAULT NULL,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_legalizes_certificate_id` (`certificate_id`),
  KEY `idx_legalizes_status` (`status`),
  CONSTRAINT `fk_legalizes_certificate` FOREIGN KEY (`certificate_id`) REFERENCES `certificates` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Question groups table
-- ============================================
CREATE TABLE IF NOT EXISTS `question_groups` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL DEFAULT '',
  `addressed_to` INT UNSIGNED NOT NULL DEFAULT 0,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Questions table
-- ============================================
CREATE TABLE IF NOT EXISTS `questions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `question_group_id` BIGINT UNSIGNED NOT NULL,
  `title` TEXT NOT NULL,
  `question_type` INT UNSIGNED NOT NULL DEFAULT 0,
  `minimum_value` VARCHAR(255) DEFAULT NULL,
  `maximum_value` VARCHAR(255) DEFAULT NULL,
  `is_required` TINYINT(1) NOT NULL DEFAULT 0,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_questions_group_id` (`question_group_id`),
  CONSTRAINT `fk_questions_group` FOREIGN KEY (`question_group_id`) REFERENCES `question_groups` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Question options table
-- ============================================
CREATE TABLE IF NOT EXISTS `question_options` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `question_id` BIGINT UNSIGNED NOT NULL,
  `title` VARCHAR(255) NOT NULL DEFAULT '',
  `is_need_essay` TINYINT(1) NOT NULL DEFAULT 0,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_question_options_question_id` (`question_id`),
  CONSTRAINT `fk_question_options_question` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Tracers table (questionnaire submissions)
-- ============================================
CREATE TABLE IF NOT EXISTS `tracers` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_tracers_user_id` (`user_id`),
  CONSTRAINT `fk_tracers_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- User answers table
-- ============================================
CREATE TABLE IF NOT EXISTS `user_answers` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `tracer_id` BIGINT UNSIGNED NOT NULL,
  `question_id` BIGINT UNSIGNED NOT NULL,
  `answer` TEXT NOT NULL,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_answers_tracer_question` (`tracer_id`, `question_id`),
  KEY `idx_user_answers_tracer_id` (`tracer_id`),
  KEY `idx_user_answers_question_id` (`question_id`),
  CONSTRAINT `fk_user_answers_tracer` FOREIGN KEY (`tracer_id`) REFERENCES `tracers` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_answers_question` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Request passwords table (password reset tokens)
-- ============================================
CREATE TABLE IF NOT EXISTS `request_passwords` (
  `id` VARCHAR(36) NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `is_used` TINYINT(1) NOT NULL DEFAULT 0,
  `expired_at` DATETIME NOT NULL,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_request_passwords_user_id` (`user_id`),
  CONSTRAINT `fk_request_passwords_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Seed users (all roles)
-- ============================================

-- Admin (user_type=3) - password: admin123
INSERT INTO `users` (`email`, `password`, `name`, `user_type`, `is_actived`) VALUES
('admin@poltekkes-medan.ac.id', '$2b$12$0FTlWutfZ6Ge.uiXyG8ZH.cLRCs49dH79XN01x2HYYc/ZAzVOLofi', 'Administrator', 3, 1);

-- Pejabat (user_type=4) - password: pejabat123
INSERT INTO `users` (`email`, `password`, `name`, `user_type`, `is_actived`) VALUES
('pejabat@poltekkes-medan.ac.id', '$2b$12$XJ2ytD7ebtJRXwmhMuKyouK1CFiNqr/9fALzYQwxhVpYg2rB9bzQ.', 'Pejabat Penandatangan', 4, 1);

-- Alumni (user_type=1) - password: alumni123
INSERT INTO `users` (`email`, `password`, `name`, `user_type`, `is_actived`) VALUES
('alumni@poltekkes-medan.ac.id', '$2b$12$L6pGwWn5SCZ2sitly6stzehTPkAzbBqsFbq0DvZ41N0NowqBlsJdq', 'Budi Santoso', 1, 1);

-- Appraiser (user_type=2) - password: appraiser123
INSERT INTO `users` (`email`, `password`, `name`, `user_type`, `is_actived`) VALUES
('appraiser@poltekkes-medan.ac.id', '$2b$12$qeidhUe3AphenNu.qyVoNulM4z7.daevDrSU1xF0iuBDSNS.QVW6u', 'Dr. Siti Aminah', 2, 1);

-- ============================================
-- Seed alumni data
-- ============================================
INSERT INTO `alumni` (`user_id`, `name`, `nik`, `place_of_birth`, `date_of_birth`, `phone`) VALUES
(3, 'Budi Santoso', '1234567890123456', 'Medan', '1995-05-15', '081234567890');

-- ============================================
-- Seed certificate
-- ============================================
INSERT INTO `certificates` (`alumni_id`, `nim`, `no_ijazah`, `major_study`, `entry_year`, `graduation_year`) VALUES
(1, '2018001001', 'Ijazah-2022-001', 'Keperawatan', '2018', '2022');

-- ============================================
-- Seed question groups for tracer study
-- ============================================
INSERT INTO `question_groups` (`id`, `title`, `addressed_to`) VALUES
(1, 'Kondisi Setelah Lulus', 1),
(2, 'Kesesuaian Kerja', 1),
(3, 'Penilaian Kompetensi', 2);

-- ============================================
-- Seed questions
-- ============================================
INSERT INTO `questions` (`id`, `question_group_id`, `title`, `question_type`, `minimum_value`, `maximum_value`, `is_required`) VALUES
-- Group 1: Kondisi Setelah Lulus
(1, 1, 'Apakah Anda sudah bekerja?', 2, NULL, NULL, 1),
(2, 1, 'Berapa lama waktu yang dibutuhkan untuk mendapatkan pekerjaan pertama?', 2, NULL, NULL, 1),
(3, 1, 'Apakah pekerjaan Anda sesuai dengan bidang studi?', 2, NULL, NULL, 1),
-- Group 2: Kesesuaian Kerja
(4, 2, 'Nama perusahaan/instansi tempat bekerja', 1, NULL, NULL, 1),
(5, 2, 'Jabatan saat ini', 1, NULL, NULL, 1),
(6, 2, 'Bidang pekerjaan', 2, NULL, NULL, 1),
-- Group 3: Penilaian Kompetensi
(7, 3, 'Bagaimana penilaian kompetensi alumni?', 2, '1', '5', 1),
(8, 3, 'Apakah alumni memiliki kemampuan komunikasi yang baik?', 2, NULL, NULL, 1);

-- ============================================
-- Seed question options
-- ============================================
INSERT INTO `question_options` (`id`, `question_id`, `title`, `is_need_essay`) VALUES
-- Q1: Apakah sudah bekerja?
(1, 1, 'Ya', 0),
(2, 1, 'Tidak', 0),
-- Q2: Waktu mendapat kerja
(3, 2, 'Kurang dari 3 bulan', 0),
(4, 2, '3-6 bulan', 0),
(5, 2, '6-12 bulan', 0),
(6, 2, 'Lebih dari 1 tahun', 0),
-- Q3: Sesuai bidang studi?
(7, 3, 'Sangat sesuai', 0),
(8, 3, 'Sesuai', 0),
(9, 3, 'Kurang sesuai', 0),
(10, 3, 'Tidak sesuai', 0),
-- Q6: Bidang pekerjaan
(11, 6, 'Kesehatan', 0),
(12, 6, 'Pendidikan', 0),
(13, 6, 'Pemerintahan', 0),
(14, 6, 'Swasta', 0),
(15, 6, 'Wirausaha', 0),
-- Q7: Penilaian kompetensi (1-5 scale)
(20, 7, '1', 0),
(21, 7, '2', 0),
(22, 7, '3', 0),
(23, 7, '4', 0),
(24, 7, '5', 0),
-- Q8: Kemampuan komunikasi
(16, 8, 'Sangat baik', 0),
(17, 8, 'Baik', 0),
(18, 8, 'Cukup', 0),
(19, 8, 'Kurang', 0);
