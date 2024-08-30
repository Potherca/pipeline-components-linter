package repositoryContents

import "time"

type LogEntry struct {
	//Author    string
	//Committer string
	//Encoding MessageEncoding
	//Hash      string
	//Message   string
	//ParentHashes []string
	//PGPSignature string
	Timestamp time.Time
}

type Logs []LogEntry

func (l Logs) Len() int {
	return len(l)
}

func (l Logs) First() LogEntry {
	return l[0]
}

func (l Logs) Last() LogEntry {
	return l[len(l)-1]
}
