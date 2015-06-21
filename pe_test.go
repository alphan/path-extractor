package main

import "testing"

func TestGitIgnore(t *testing.T) {
	output := getAllMatches("?? alt/generateStore.php")
	if output[0] != "alt/generateStore.php" {
		t.Errorf("Doesnt match files", output[0])
	}

	output = getAllMatches("hello .gitignore")
	if output[0] != ".gitignore" {
		t.Errorf("Doesnt match hidden files", output[0])
	}

	output = getAllMatches(" mail@mail.com ")
	if len(output) != 0 {
		t.Errorf("Matches email adresses", output[0])
	}

	output = getAllMatches("v1.2")
	if len(output) != 0 {
		t.Errorf("Matches version number", output[0])
	}

	output = getAllMatches("obj.slice()")
	if len(output) != 0 {
		t.Errorf("Matches function call", output[0])
	}

	output = getAllMatches("~/www")
	if len(output) == 0 || output[0] != "~/www" {
		t.Errorf("Doesnt match home", output[0])
	}

	output = getAllMatches("origin/master")
	if len(output) != 0 {
		t.Errorf("Matches remote name", output[0])
	}

	output = getAllMatches("john doe (dead on 28/04/2014)")
	if len(output) != 0 {
		t.Errorf("Matches date", output[0])
	}

	output = getAllMatches("john doe ,dead on 28/04/2014")
	if len(output) != 0 {
		t.Errorf("Matches date", output[0])
	}

	output = getAllMatches(".gitignore , ~/www")
	if len(output) != 2 {
		t.Errorf("Doesnt match multi", output[0])
	}

	output = getAllMatches("var/")
	if len(output) != 1 {
		t.Errorf("Doesnt match dir", output[0])
	}

	output = getAllMatches("//")
	if len(output) != 0 {
		t.Errorf("Comment matches", output[0])
	}
}
