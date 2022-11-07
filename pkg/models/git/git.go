package git

type AuthInfo struct {
	RemoteUrl string `json:"remoteurl" description:"git server url"`
}

type GitVerifier interface {
	VerifyGitCredential(remoteUrl string, namespace string, secretName string) error
}
