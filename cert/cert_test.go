package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid? err=%bv", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got nil")
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "qmsdhfqdfqdfhqsdkjhgfqsdiqhjdsfhqsdljkfhqsdlkfhqsdlkfhqsdlkjfhqldsjqdsfbhqlksjdfqdfqf"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be return on a too long course (course=%s)", course)
	}
}
