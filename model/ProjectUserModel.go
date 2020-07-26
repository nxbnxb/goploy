package model

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
)

const projectUserTable = "`project_user`"

// ProjectUser project user relationship
type ProjectUser struct {
	ID          int64  `json:"id"`
	ProjectID   int64  `json:"projectId"`
	ProjectName string `json:"projectName"`
	UserID      int64  `json:"userId"`
	UserName    string `json:"userName"`
	InsertTime  string `json:"insertTime"`
	UpdateTime  string `json:"updateTime"`
}

// ProjectUsers project user relationship
type ProjectUsers []ProjectUser

// GetBindUserListByProjectID user row
func (pu ProjectUser) GetBindUserListByProjectID() (ProjectUsers, error) {
	rows, err := sq.
		Select("project_user.id, project_id, user_id, user.name, project_user.insert_time, project_user.update_time").
		From(projectUserTable).
		LeftJoin(userTable + " ON project_user.user_id = user.id").
		Where(sq.Eq{"project_id": pu.ProjectID}).
		RunWith(DB).
		Query()
	if err != nil {
		return nil, err
	}
	projectUsers := ProjectUsers{}
	for rows.Next() {
		var projectUser ProjectUser

		if err := rows.Scan(&projectUser.ID, &projectUser.ProjectID, &projectUser.UserID, &projectUser.UserName, &projectUser.InsertTime, &projectUser.UpdateTime); err != nil {
			return nil, err
		}
		projectUsers = append(projectUsers, projectUser)
	}
	return projectUsers, nil
}

// GetBindUserListByProjectID user row
func (pu ProjectUser) GetBindProjectListByUserID() (ProjectUsers, error) {
	rows, err := sq.
		Select("project_user.id, project_id, user_id, project.name, project_user.insert_time, project_user.update_time").
		From(projectUserTable).
		LeftJoin(projectTable + " ON project_user.project_id = project.id").
		Where(sq.Eq{"user_id": pu.UserID}).
		RunWith(DB).
		Query()
	if err != nil {
		return nil, err
	}
	projectUsers := ProjectUsers{}
	for rows.Next() {
		var projectUser ProjectUser

		if err := rows.Scan(&projectUser.ID, &projectUser.ProjectID, &projectUser.UserID, &projectUser.ProjectName, &projectUser.InsertTime, &projectUser.UpdateTime); err != nil {
			return nil, err
		}
		projectUsers = append(projectUsers, projectUser)
	}
	return projectUsers, nil
}

// GetListByUserID user row
func (pu ProjectUser) GetListByUserID() (ProjectUsers, error) {
	rows, err := sq.
		Select("id, project_id, user_id, insert_time, update_time").
		From(projectUserTable).
		Where(sq.Eq{"user_id": pu.UserID}).
		RunWith(DB).
		Query()
	if err != nil {
		return nil, err
	}
	projectUsers := ProjectUsers{}
	for rows.Next() {
		var projectUser ProjectUser

		if err := rows.Scan(&projectUser.ID, &projectUser.ProjectID, &projectUser.UserID, &projectUser.InsertTime, &projectUser.UpdateTime); err != nil {
			return nil, err
		}
		projectUsers = append(projectUsers, projectUser)
	}
	return projectUsers, nil
}

// AddMany add many row to table project_user
func (pu ProjectUsers) AddMany() error {
	if len(pu) == 0 {
		return nil
	}
	builder := sq.
		Replace(projectUserTable).
		Columns("project_id", "user_id")

	for _, row := range pu {
		builder = builder.Values(row.ProjectID, row.UserID)
	}
	_, err := builder.RunWith(DB).Exec()
	return err
}

// AddAdminByUserID add admin to table project_user
func (pu ProjectUser) AddAdminByUserID() error {
	builder := sq.
		Replace(projectUserTable).
		Columns("project_id", "user_id").
		Select(sq.
			Select(fmt.Sprintf("id as project_id, %d as user_id", pu.UserID)).
			From(projectTable))
	_, err := builder.RunWith(DB).Exec()
	return err
}

// AddNamespaceProjectInUserID add all project with namespace in user id
func (pu ProjectUser) AddNamespaceProjectInUserID(namespaceID int64, userIDs []int64) error {
	userIDSQL := ""
	for _, userID := range userIDs {
		userIDSQL += fmt.Sprintf("SELECT %d as user_id UNION ", userID)
	}

	userIDSQL = "(" + strings.TrimRight(userIDSQL, " UNION ") + ") as t1"

	builder := sq.
		Replace(projectUserTable).
		Columns("project_id", "user_id").
		Select(sq.
			Select("project.id as project_id, t1.user_id as user_id").
			From(userIDSQL).
			Join(projectTable).
			Where(sq.Eq{"namespace_id": namespaceID}))
	_, err := builder.RunWith(DB).Exec()
	return err
}

// DeleteRow edit one row to table ProjectUser
func (pu ProjectUser) DeleteRow() error {
	_, err := sq.
		Delete(projectUserTable).
		Where(sq.Eq{"id": pu.ID}).
		RunWith(DB).
		Exec()
	return err
}

// DeleteRow edit one row to table ProjectUser
func (pu ProjectUser) DeleteByUserID() error {
	_, err := sq.
		Delete(projectUserTable).
		Where(sq.Eq{"user_id": pu.UserID}).
		RunWith(DB).
		Exec()
	return err
}
