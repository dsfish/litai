package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListFiles(t *testing.T) {
	files, err := listFiles(textDir)
	require.NoError(t, err)
	assert.NotEmpty(t, files)
}

func TestReadFile(t *testing.T) {
	files, err := listFiles(textDir)
	require.NoError(t, err)
	require.NotEmpty(t, files)

	contents, err := readFile(files[0])
	require.NoError(t, err)
	require.NotEmpty(t, contents)
}
