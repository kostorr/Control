package repos

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"strings"
)

type Repo struct {
	HostingSite string
	User string
	RepoName string
	Revision string
	Default bool
	//Properties RepoProperties
}

/*type RepoProperties struct {
	Default bool
	Priority int
}*/

func NewRepo(repoPath string) (*Repo, error){

	revSlice := strings.Split(repoPath, "@")

	var repoUrlSlice []string
	var revision string
	if len(revSlice) == 2 { //revision specified
		repoUrlSlice = strings.Split(revSlice[0], "/")
		revision = revSlice[1]
	} else if len(revSlice) == 1{ //no revision specified
		repoUrlSlice = strings.Split(revSlice[0], "/")
		revision = "" // TODO: or master...
	} else {
		return &Repo{}, errors.New("Repo path resolution failed")
	}

	return &Repo{repoUrlSlice[0], repoUrlSlice[1],
		repoUrlSlice[2], revision, false}, nil
}

func (r *Repo) GetIdentifier() string {
	identifier := r.HostingSite + "/" + r.User + "/" + r.RepoName

	return identifier
}

func (r *Repo) GetCompleteIdentifier() string {
	identifier := r.HostingSite + "/" + r.User + "/" + r.RepoName

	if r.Revision != "" {
		identifier += "@" + r.Revision
	}

	return identifier
}

func (r *Repo) GetCloneDir() string {
	cloneDir := viper.GetString("repositoriesUri")
	if cloneDir[len(cloneDir)-1:] != "/" {
		cloneDir += "/"
	}

	cloneDir += r.HostingSite + "/" +
				r.User 		 + "/" +
				r.RepoName

	return cloneDir
}

func (r *Repo) GetUrl() string {
	return "https://" +
		r.HostingSite + "/" +
		r.User 		  + "/" +
		r.RepoName
}

func (r *Repo) GetTaskDir() string {
	return r.GetCloneDir() + "/tasks/"
}

func (r *Repo) GetWorkflowDir() string {
	return r.GetCloneDir() + "/workflows/"
}

func (r *Repo) ResolveTaskClassIdentifier(loadTaskClass string) string {
		taskClassIdentifier := r.HostingSite + "/" + r.User + "/" + r.RepoName + "/" + loadTaskClass

	if r.Revision != "" {
		taskClassIdentifier += "@" + r.Revision
	}

	return taskClassIdentifier
}

func (r *Repo) checkoutBranch(branch string) error {
	if branch == "" {
		branch = "master"
	}

	ref, err := git.PlainOpen(r.GetCloneDir())
	if err != nil {
		return err
	}

	w, err := ref.Worktree()
	if err != nil {
		return err
	}

	checkErr := w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})

	if checkErr != nil {
		return err
	}

	return nil
}
