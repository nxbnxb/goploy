package model

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

const projectTable = "`project`"

// Project mysql table project
type Project struct {
	ID                    int64  `json:"id"`
	GroupID               int64  `json:"groupId"`
	Name                  string `json:"name"`
	URL                   string `json:"url"`
	Path                  string `json:"path"`
	SymlinkPath           string `json:"symlinkPath"`
	Environment           string `json:"environment"`
	Branch                string `json:"branch"`
	AfterPullScriptMode   string `json:"afterPullScriptMode"`
	AfterPullScript       string `json:"afterPullScript"`
	AfterDeployScriptMode string `json:"afterDeployScriptMode"`
	AfterDeployScript     string `json:"afterDeployScript"`
	RsyncOption           string `json:"rsyncOption"`
	AutoDeploy            uint8  `json:"autoDeploy"`
	PublisherID           int64  `json:"publisherId"`
	PublisherName         string `json:"publisherName"`
	PublishExt            string `json:"publishExt"`
	DeployState           uint8  `json:"deployState"`
	LastPublishToken      string `json:"lastPublishToken"`
	NotifyType            uint8  `json:"notifyType"`
	NotifyTarget          string `json:"notifyTarget"`
	State                 uint8  `json:"state"`
	InsertTime            string `json:"insertTime"`
	UpdateTime            string `json:"updateTime"`
}

const (
	ProjectNotDeploy = 0
	ProjectDeploying = 1
	ProjectSuccess   = 2
	ProjectFail      = 3
)

const (
	ProjectManualDeploy  = 0
	ProjectWebhookDeploy = 1
)

// Projects many project
type Projects []Project

// AddRow add one row to table project and add id to p.ID
func (p Project) AddRow() (int64, error) {
	result, err := sq.
		Insert(projectTable).
		Columns("group_id", "name", "url", "path", "symlink_path", "environment", "branch", "after_pull_script_mode", "after_pull_script", "after_deploy_script_mode", "after_deploy_script", "rsync_option", "auto_deploy", "notify_type", "notify_target").
		Values(p.GroupID, p.Name, p.URL, p.Path, p.SymlinkPath, p.Environment, p.Branch, p.AfterPullScriptMode, p.AfterPullScript, p.AfterDeployScriptMode, p.AfterDeployScript, p.RsyncOption, p.AutoDeploy, p.NotifyType, p.NotifyTarget).
		RunWith(DB).
		Exec()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

// EditRow edit one row to table project
func (p Project) EditRow() error {
	_, err := sq.
		Update(projectTable).
		SetMap(sq.Eq{
			"group_id":                 p.GroupID,
			"name":                     p.Name,
			"url":                      p.URL,
			"path":                     p.Path,
			"symlink_path":             p.SymlinkPath,
			"environment":              p.Environment,
			"branch":                   p.Branch,
			"after_pull_script_mode":   p.AfterPullScriptMode,
			"after_pull_script":        p.AfterPullScript,
			"after_deploy_script_mode": p.AfterDeployScriptMode,
			"after_deploy_script":      p.AfterDeployScript,
			"rsync_option":             p.RsyncOption,
			"auto_deploy":              p.AutoDeploy,
			"notify_type":              p.NotifyType,
			"notify_target":            p.NotifyTarget,
		}).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		Exec()
	return err
}

// RemoveRow project
func (p Project) RemoveRow() error {
	_, err := sq.
		Update(projectTable).
		SetMap(sq.Eq{
			"state": Disable,
		}).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		Exec()
	return err
}

// Publish for project
func (p Project) Publish() error {
	_, err := sq.
		Update(projectTable).
		SetMap(sq.Eq{
			"publisher_id":       p.PublisherID,
			"publisher_name":     p.PublisherName,
			"deploy_state":       p.DeployState,
			"last_publish_token": p.LastPublishToken,
		}).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		Exec()
	return err
}

// DeploySuccess return err
func (p Project) DeploySuccess() error {
	_, err := sq.
		Update(projectTable).
		SetMap(sq.Eq{
			"deploy_state": ProjectSuccess,
		}).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		Exec()

	return err
}

// DeploySuccess return err
func (p Project) DeployFail() error {
	_, err := sq.
		Update(projectTable).
		SetMap(sq.Eq{
			"deploy_state": ProjectFail,
		}).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		Exec()

	return err
}

// GetList project row
func (p Project) GetList(pagination Pagination, groupIDs []string) (Projects, error) {
	builder := sq.
		Select("id, group_id, name, url, path, symlink_path, environment, branch, after_pull_script_mode, after_pull_script, after_deploy_script_mode, after_deploy_script, rsync_option, auto_deploy, notify_type, notify_target, insert_time, update_time").
		From(projectTable).
		Where(sq.Eq{"state": Enable})

	if len(p.Name) > 0 {
		builder = builder.Where(sq.Like{"name": "%" + p.Name + "%"})
	}

	if len(groupIDs) > 0 {
		builder = builder.Where(sq.Eq{"group_id": groupIDs})
	}

	rows, err := builder.Limit(pagination.Rows).
		Offset((pagination.Page - 1) * pagination.Rows).
		OrderBy("id DESC").
		RunWith(DB).
		Query()

	if err != nil {
		return nil, err
	}
	projects := Projects{}
	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.SymlinkPath,
			&project.Environment,
			&project.Branch,
			&project.AfterPullScriptMode,
			&project.AfterPullScript,
			&project.AfterDeployScriptMode,
			&project.AfterDeployScript,
			&project.RsyncOption,
			&project.AutoDeploy,
			&project.NotifyType,
			&project.NotifyTarget,
			&project.InsertTime,
			&project.UpdateTime,
		); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

// GetList project total
func (p Project) GetTotal(groupIDs []string) (int64, error) {
	var total int64
	builder := sq.
		Select("COUNT(*) AS count").
		From(projectTable).
		Where(sq.Eq{"state": Enable})
	if len(p.Name) > 0 {
		builder = builder.Where(sq.Like{"name": "%" + p.Name + "%"})
	}
	if len(groupIDs) > 0 {
		builder = builder.Where(sq.Eq{"group_id": groupIDs})
	}
	err := builder.RunWith(DB).
		QueryRow().
		Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// GetListByManagerGroupStr project row
func (p Project) GetListInGroupIDs(groupIDs []string, pagination Pagination) (Projects, Pagination, error) {
	builder := sq.
		Select("id, group_id, name, url, path, symlink_path, environment, branch, after_pull_script_mode, after_pull_script, after_deploy_script_mode, after_deploy_script, rsync_option, auto_deploy, notify_type, notify_target, insert_time, update_time").
		From(projectTable).
		Where(sq.Eq{"group_id": groupIDs}).
		Where(sq.Like{"name": "%" + p.Name + "%"}).
		Where(sq.Eq{"state": Enable}).
		Limit(pagination.Rows).
		Offset((pagination.Page - 1) * pagination.Rows).
		OrderBy("id DESC")
	pageBuilder := sq.
		Select("COUNT(*) AS count").
		Where(sq.Eq{"group_id": groupIDs}).
		Where(sq.Like{"name": "%" + p.Name + "%"}).
		Where(sq.Eq{"state": Enable}).
		From(projectTable)
	rows, err := builder.RunWith(DB).Query()
	if err != nil {
		return nil, pagination, err
	}
	projects := Projects{}
	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.SymlinkPath,
			&project.Environment,
			&project.Branch,
			&project.AfterPullScriptMode,
			&project.AfterPullScript,
			&project.AfterDeployScriptMode,
			&project.AfterDeployScript,
			&project.RsyncOption,
			&project.AutoDeploy,
			&project.NotifyType,
			&project.NotifyTarget,
			&project.InsertTime,
			&project.UpdateTime,
		); err != nil {
			return nil, pagination, err
		}
		projects = append(projects, project)
	}
	err = pageBuilder.
		RunWith(DB).
		QueryRow().
		Scan(&pagination.Total)
	if err != nil {
		return nil, pagination, err
	}
	return projects, pagination, nil
}

func (p Project) GetUserProjectList(userID int64, userRole string, groupIDStr string) (Projects, error) {
	builder := sq.
		Select(`
			project.id, 
			project.name,
			project.publisher_id,
			project.publisher_name,
			IFNULL(publish_trace.ext, '{}'),
			project.group_id,
			project.environment, 
			project.branch, 
			project.last_publish_token,
			project.deploy_state, 
			project.update_time`).
		Where(sq.Eq{"project.state": Enable}).
		OrderBy("project.id DESC")
	if userRole == "admin" || userRole == "manager" {
		builder = builder.From(projectTable).
			LeftJoin(fmt.Sprintf("%[1]s ON %[1]s.token = %s.last_publish_token and type = %d", publishTraceTable, projectTable, Pull))
	} else if userRole == "group-manager" {
		builder = builder.From(projectTable).
			LeftJoin(fmt.Sprintf("%[1]s ON %[1]s.token = %s.last_publish_token and type = %d", publishTraceTable, projectTable, Pull)).
			Where("group_id IN (?) or project.id in (select project_id from "+projectUserTable+" where user_id = ?)", groupIDStr, userID)
	} else {
		builder = builder.From(projectUserTable).
			LeftJoin(projectTable + " ON project_user.project_id = project.id").
			LeftJoin(fmt.Sprintf("%[1]s ON %[1]s.token = %s.last_publish_token and type = %d", publishTraceTable, projectTable, Pull)).
			Where(sq.Eq{"project_user.user_id": userID})
	}

	if p.GroupID != 0 {
		builder = builder.Where(sq.Eq{"project.group_id": p.GroupID})
	}
	if len(p.Name) > 0 {
		builder = builder.Where(sq.Like{"project.name": "%" + p.Name + "%"})
	}

	rows, err := builder.RunWith(DB).Query()
	if err != nil {
		return nil, err
	}
	projects := Projects{}
	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.PublisherID,
			&project.PublisherName,
			&project.PublishExt,
			&project.GroupID,
			&project.Environment,
			&project.Branch,
			&project.LastPublishToken,
			&project.DeployState,
			&project.UpdateTime); err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	return projects, nil

}

// GetAll Group row
func (p Project) GetAll() (Projects, error) {
	rows, err := sq.
		Select("id, group_id, name, url, path, environment, branch, rsync_option, deploy_state").
		From(projectTable).
		Where(sq.Eq{"state": Enable}).
		OrderBy("id DESC").
		RunWith(DB).
		Query()
	if err != nil {
		return nil, err
	}
	projects := Projects{}
	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.Environment,
			&project.Branch,
			&project.RsyncOption,
			&project.DeployState); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// GetAll Group row
func (p Project) GetAllByName() (Projects, error) {
	rows, err := sq.
		Select("id, group_id, name, url, path, environment, branch, rsync_option, deploy_state").
		From(projectTable).
		Where(sq.Eq{"name": p.Name}).
		Where(sq.Eq{"state": Enable}).
		OrderBy("id DESC").
		RunWith(DB).
		Query()
	if err != nil {
		return nil, err
	}
	projects := Projects{}
	for rows.Next() {
		var project Project

		if err := rows.Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.Environment,
			&project.Branch,
			&project.RsyncOption,
			&project.DeployState); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// GetData add project information to p *Project
func (p Project) GetData() (Project, error) {
	var project Project
	err := sq.
		Select("id, group_id, name, url, path, symlink_path, environment, branch, after_pull_script_mode, after_pull_script, after_deploy_script_mode, after_deploy_script, rsync_option, auto_deploy, deploy_state, notify_type, notify_target, insert_time, update_time").
		From(projectTable).
		Where(sq.Eq{"id": p.ID}).
		RunWith(DB).
		QueryRow().
		Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.SymlinkPath,
			&project.Environment,
			&project.Branch,
			&project.AfterPullScriptMode,
			&project.AfterPullScript,
			&project.AfterDeployScriptMode,
			&project.AfterDeployScript,
			&project.RsyncOption,
			&project.AutoDeploy,
			&project.DeployState,
			&project.NotifyType,
			&project.NotifyTarget,
			&project.InsertTime,
			&project.UpdateTime)
	if err != nil {
		return project, err
	}
	return project, nil
}

// GetData add project information to p *Project
func (p Project) GetDataByName() (Project, error) {
	var project Project
	err := sq.
		Select("id, group_id, name, url, path, symlink_path, environment, branch, after_pull_script_mode, after_pull_script, after_deploy_script_mode, after_deploy_script, rsync_option, auto_deploy, deploy_state, notify_type, notify_target, insert_time, update_time").
		From(projectTable).
		Where(sq.Eq{"name": p.Name}).
		RunWith(DB).
		QueryRow().
		Scan(
			&project.ID,
			&project.GroupID,
			&project.Name,
			&project.URL,
			&project.Path,
			&project.SymlinkPath,
			&project.Environment,
			&project.Branch,
			&project.AfterPullScriptMode,
			&project.AfterPullScript,
			&project.AfterDeployScriptMode,
			&project.AfterDeployScript,
			&project.RsyncOption,
			&project.AutoDeploy,
			&project.DeployState,
			&project.NotifyType,
			&project.NotifyTarget,
			&project.InsertTime,
			&project.UpdateTime)
	if err != nil {
		return project, err
	}
	return project, nil
}

func (p Project) GetUserProjectData(userID int64, userRole string, groupIDStr string) (Project, error) {
	var builder sq.SelectBuilder
	if userRole == "admin" || userRole == "manager" {
		builder = sq.
			Select("id, name, publisher_id, publisher_name, group_id, environment, branch, last_publish_token, deploy_state, update_time").
			From(projectTable).
			Where(sq.Eq{"id": p.ID}).
			Where(sq.Eq{"state": Enable})
	} else if userRole == "group-manager" {
		builder = sq.
			Select("id, name, publisher_id, publisher_name, group_id, environment, branch, last_publish_token, deploy_state, update_time").
			From(projectTable).
			Where("group_id IN (?) or id in (select project_id from "+projectUserTable+" where user_id = ?)", groupIDStr, userID).
			Where(sq.Eq{"id": p.ID}).
			Where(sq.Eq{"state": Enable})
	} else {
		builder = sq.
			Select("project_id, project.name, publisher_id, publisher_name, project.group_id, project.environment, project.branch, project.last_publish_token, project.deploy_state, project.update_time").
			From(projectUserTable).
			LeftJoin(projectTable + " ON project_user.project_id = project.id").
			Where(sq.Eq{
				"project_user.project_id": p.ID,
				"project_user.user_id":    userID,
				"project.state":           Enable,
			})
	}
	var project Project
	err := builder.
		RunWith(DB).
		QueryRow().
		Scan(&project.ID,
			&project.Name,
			&project.PublisherID,
			&project.PublisherName,
			&project.GroupID,
			&project.Environment,
			&project.Branch,
			&project.LastPublishToken,
			&project.DeployState,
			&project.UpdateTime)
	if err != nil {
		return project, err
	}
	return project, nil
}
