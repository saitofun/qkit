// This is a generated source file. DO NOT EDIT
// Version: 0.0.1
// Source: must/must.go
// Date: Jul  8 01:13:34

package must

import "log"

// pls add your assert function here, or add type and re-generate

func Must(err error) {
	if err != nil {
		log.Panic(err)
	}
}
