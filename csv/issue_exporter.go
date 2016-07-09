package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/b0m0x/gitlab-issue-exporter/gitlab"
	"io"
)

type CsvIssueWriter struct {
	writer *csv.Writer
}

func Priority(label []string) string {
	for i := 0; i < len(label); i++ {
		if label[i] == "一般程度" {
			return "Low"
		}

		if label[i] == "严重程度" {
			return "Normal"
		}

		if label[i] == "高优先级" {
			return "High"
		}

		if label[i] == "紧急处理" {
			return "Urgent"
		}
	}

	return "Normal"
}

func Class(label []string) string {
	for i := 0; i < len(label); i++ {
		if label[i] == "功能bug" {
			return "功能"
		}

		if label[i] == "性能bug" {
			return "性能"
		}

		if label[i] == "体验性bug" {
			return "体验性"
		}

		if label[i] == "缺陷预防" {
			return "缺陷预防"
		}

		if label[i] == "UI-bug" {
			return "UI"
		}

		if label[i] == "需求" {
			return "需求"
		}
	}

	return "体验性"
}
func State(label []string) string {
	for i := 0; i < len(label); i++ {
		if label[i] == "new" {
			return "new"
		}

		if label[i] == "fixed_待审核" {
			return "fixed_待审核"
		}

		if label[i] == "fixed_待回归" {
			return "fixed_待回归"
		}

		if label[i] == "halt_1" {
			return "halt_1"
		}

		if label[i] == "halt_2" {
			return "halt_2"
		}
	}

	return "new"
}

func NameTrans(name string) string {

	if name == "何倩华47993" {
		return "何 倩华"
	}

	if name == "吴利斌68923" {
		return "吴 利斌"
	}

	if name == "黄腾辉71539" {
		return "黄 腾辉"
	}
	if name == "欧阳14034" {
		return "欧 阳"
	}

	if name == "叶德华70651" {
		return "叶 德华"
	}

	if name == "钟德财27927" {
		return "钟 德财"
	}

	if name == "黄茂彪45637" {
		return "黄 茂彪"
	}
	if name == "刘少东16700" {
		return "刘 少东"
	}
	return "吴 利斌"
}

func (w *CsvIssueWriter) Write(issue *gitlab.GitlabIssue) {
	w.writer.Write([]string{
		fmt.Sprintf("%d", issue.Id),
		Priority(issue.Labels),
		Class(issue.Labels),
		State(issue.Labels),
		issue.Title,
		issue.Description,
		issue.CreatedAt.Format("02/01/2006"),
		issue.Milestone.Title,
		NameTrans(issue.Assignee.Name),
		NameTrans(issue.Author.Name),
	})
	w.writer.Flush()
}

func NewCsvIssueWriter(w io.Writer) *CsvIssueWriter {
	cw := &CsvIssueWriter{csv.NewWriter(w)}
	cw.writer.Write([]string{
		"id",
		"priority",
		"class",
		"state",
		"title",
		"description",
		"create_at",
		"milestone",
		"assigned_to",
		"author",
	})
	return cw
}
