package main

import (
	"testing"
	
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) { goleak.VerifyTestMain(m)}
func Test_main(t *testing.T) { main() }