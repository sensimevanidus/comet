package rest

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var internalStorage map[string]*storage

// storage is an in-memory key-value store that is populated after a
// configuration file is read. It is used to store variables provided in the
// configuration file along with environment variables, and values meant to be
// stored from the response of API calls. Note that there's a cascading mechanism
// for the variable keys. This means that if a key is present both as an
// environment variable and a definition in the configuration file, the
// configuration file overrides the other one.
type storage struct {
	data  map[string]string
	mutex *sync.RWMutex
}

func initStorage(conf *testConfiguration) {
	if internalStorage == nil {
		internalStorage = make(map[string]*storage, 0)
	}

	internalStorage[conf.Name] = &storage{
		data:  conf.Variables,
		mutex: &sync.RWMutex{},
	}

	internalStorage[conf.Name].enrichStorageWithEnvironmentVariables()
}

func getStorage(testName string) *storage {
	if internalStorage == nil {
		return nil
	}

	if s, ok := internalStorage[testName]; ok {
		return s
	}

	return nil
}

func (s *storage) write(fieldKey, fieldValue string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[fieldKey] = fieldValue
	return
}

func (s *storage) read(fieldKey string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if value, ok := s.data[fieldKey]; ok {
		return value, nil
	}

	return "", fmt.Errorf("%v not found in storage", fieldKey)
}

func (s *storage) enrichStorageWithEnvironmentVariables() {
	for _, e := range os.Environ() {
		envVariable := strings.SplitN(e, "=", 2)
		if v, err := s.read(envVariable[0]); err != nil && v == "" {
			s.write(envVariable[0], envVariable[1])
		}
	}
}
