package model

import (
	"context"
	"database/sql"
	"strings"
	"tracert/internal/pkg/util"
	"tracert/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlumniAppraiser struct {
	Pb proto.AlumniAppraiser
}

func (u *AlumniAppraiser) Create(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return util.ContextError(ctx)
	default:
	}

	query := `INSERT INTO alumni_appraisers (user_id, alumni_id, name, instansi, position, alumni_position) VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare insert: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		u.Pb.UserId,
		u.Pb.Alumni.Id,
		u.Pb.Name,
		u.Pb.Instansi,
		u.Pb.Position,
		u.Pb.AlumniPosition,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec insert: %v", err)
	}

	return nil
}

func (u *AlumniAppraiser) ListQuery(ctx context.Context, db *sql.DB, in *proto.ListInput) (string, []interface{}, *proto.ListInput, error) {
	query := `SELECT id, user_id, name, instansi, position, alumni_id, alumni_position, created, modified FROM alumni_appraisers`
	where := []string{}
	paramQueries := []interface{}{}

	if len(in.Search) > 0 {
		paramQueries = append(paramQueries, in.Search)
		paramQueries = append(paramQueries, in.Search)
		paramQueries = append(paramQueries, in.Search)
		where = append(where, `(name LIKE ? OR instansi LIKE ? OR position LIKE ?)`)
	}

	{
		qCount := `SELECT COUNT(*) FROM alumni_appraisers`
		if len(where) > 0 {
			qCount += " WHERE " + strings.Join(where, " AND ")
		}
		var count int
		err := db.QueryRowContext(ctx, qCount, paramQueries...).Scan(&count)
		if err != nil && err != sql.ErrNoRows {
			return query, paramQueries, in, status.Error(codes.Internal, err.Error())
		}

		in.Count = uint32(count)
	}

	if len(where) > 0 {
		query += ` WHERE ` + strings.Join(where, " AND ")
	}

	if len(in.OrderBy) == 0 || !(in.OrderBy == "name" || in.OrderBy == "instansi" || in.OrderBy == "position") {
		if in == nil {
			in = &proto.ListInput{OrderBy: "created"}
		} else {
			in.OrderBy = "created"
		}
	}

	query += ` ORDER BY ` + in.OrderBy

	if len(in.Sort) > 0 {
		sort := strings.ToLower(in.Sort)
		if sort == "desc" {
			query += ` ` + sort
		} else {
			in.Sort = "asc"
		}
	}

	if in.Limit > 0 {
		query += ` LIMIT ? OFFSET ?`
		page := in.Page
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * in.Limit
		paramQueries = append(paramQueries, in.Limit, offset)
	}

	return query, paramQueries, in, nil
}
