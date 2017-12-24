package dbt

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// TestRepo a fake repository server.  Basically an in-memory http server that can be used as a test fixture for testing the internal API.  Cool huh?
type TestRepo struct{}

// Run runs the test repository server.
func (tr *TestRepo) Run(port int) (err error) {

	log.Printf("Running test artifact server on port %d", port)

	http.HandleFunc("/dbt/truststore", tr.HandlerTruststore)

	http.HandleFunc("/dbt-tools/foo/", tr.HandlerFoo)

	http.HandleFunc("/dbt-tools/foo/1.2.2/", tr.HandlerVersionA)

	http.HandleFunc("/dbt-tools/foo/1.2.3/", tr.HandlerVersionB)

	err = http.ListenAndServe(fmt.Sprintf("localhost:%s", strconv.Itoa(port)), nil)

	return err
}

// HandlerTruststore handles requests on the dbt repo path
func (tr *TestRepo) HandlerTruststore(w http.ResponseWriter, r *http.Request) {
	log.Printf("*TestRepo: DBT Request for %s*", r.URL.Path)

	_, err := w.Write([]byte(testKeyPublic()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("500 - %s", err)))
	}
}

// HandlerFoo handles requests for the 'foo' tool
func (tr *TestRepo) HandlerFoo(w http.ResponseWriter, r *http.Request) {
	log.Printf("**TestRepo: Tool Request for %s", r.URL.Path)

	switch r.URL.Path {
	// The index
	case "/dbt-tools/foo/":
		_, err := w.Write([]byte(dbtIndexOutput()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// HandlerVersionA handles requests for version A
func (tr *TestRepo) HandlerVersionA(w http.ResponseWriter, r *http.Request) {
	log.Printf("**TestRepo: Tool Request for %s", r.URL.Path)

	switch r.URL.Path {
	// Index
	case "/dbt-tools/foo/1.2.2/":
		_, err := w.Write([]byte(dbtVersionAIndexOutput()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.2/linux/x86_64/foo":
		_, err := w.Write([]byte(dbtVersionAContent()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.2/linux/x86_64/foo.md5":
		_, err := w.Write([]byte(dbtVersionAMd5()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.2/linux/x86_64/foo.sha1":
		_, err := w.Write([]byte(dbtVersionASha1()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.2/linux/x86_64/foo.sha256":
		_, err := w.Write([]byte(dbtVersionASha256()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.2/linux/x86_64/foo.asc":
		_, err := w.Write([]byte(dbtVersionASig()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// HandlerVersionB  handles requests for version B
func (tr *TestRepo) HandlerVersionB(w http.ResponseWriter, r *http.Request) {
	log.Printf("**TestRepo: Tool Request for %s", r.URL.Path)

	switch r.URL.Path {
	// The index
	case "/dbt-tools/foo/1.2.3/":
		_, err := w.Write([]byte(dbtVersionBIndexOutput()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.3/linux/x86_64/foo":
		_, err := w.Write([]byte(dbtVersionBContent()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.3/linux/x86_64/foo.md5":
		_, err := w.Write([]byte(dbtVersionBMd5()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.3/linux/x86_64/foo.sha1":
		_, err := w.Write([]byte(dbtVersionBSha1()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.3/linux/x86_64/foo.sha256":
		_, err := w.Write([]byte(dbtVersionBSha256()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	case "/dbt-tools/foo/1.2.3/linux/x86_64/foo.asc":
		_, err := w.Write([]byte(dbtVersionBSig()))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
