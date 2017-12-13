package dbt

import "fmt"

func testDbtConfigContents(port int) string {
	return fmt.Sprintf(`{
  "dbt": {
    "repository": "http://localhost:%d/dbt",
    "truststore": "http://localhost:%d/dbt/truststore"
  },
  "tools": {
    "repository": "http://localhost:%d/dbt-tools"
  }
}`, port, port, port)
}

func testDbtConfig(port int) Config {
	return Config{
		DbtConfig{
			Repo:       fmt.Sprintf("http://localhost:%d/dbt", port),
			TrustStore: fmt.Sprintf("http://localhost:%d/dbt/truststore", port),
		},
		ToolsConfig{
			Repo: fmt.Sprintf("http://localhost:%d/dbt-tools", port),
		},
	}
}

func testDbtObj(port int) *DBT {
	dbtObj := &DBT{
		Config:  testDbtConfig(port),
		Verbose: true,
	}

	return dbtObj
}

// public key for testing
func testKeyPublic() string {
	return `-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFowLigBCAC++pVrVRRM86Wo8V7XJsOmU2xtBBY5a8ktB1tdpEhzlPWQHObx
LINj79HE3lRlIFQmxnKcX3I15bzT3yo3XWLyVUsCDA1Mg9JoU2zJ+u3XftdNBg8J
eRlTiEwZYflxEYZFSyh3TZI2VZxxlINp/jOGG0dpAdKF3sfKxdTRb30lgDr+wIzv
oncrjX023UQDHoRZ3f+zPpnkubjhwH8jUHLiGsyKvu0XDB0c4y/6yG6vLUMQDuKX
bkzBtssdLLA6MTur9Q26dQV/DvuNZdHx17vwXSvf/JMKdWcX80fsAJD644KW9DOg
pgLqtBa4Tfutt3S8ueIHDnPZBKFL0u+Q61xvABEBAAG0HERCVCBUZXN0IFN1aXRl
IDxkYnRAZGJ0LmNvbT6JAVQEEwEIAD4WIQTdbDzq2B9JD2WAtKLOaEY1/aXTHwUC
WjAuKAIbAwUJA8JnAAULCQgHAgYVCAkKCwIEFgIDAQIeAQIXgAAKCRDOaEY1/aXT
H52LCACYqQnVmJRarckqh1//FUFFpXlTcwWV2zGr3CEFRs0BrWEQD7giehFpKoTL
JOJJSFd4xcbo/9wMXpJ16soK83o48laxkj+2LDUfDylnTVpVI6zVvAseqnt5nbrA
CWes75FeIHtQ6woDy7K3RHUORNZ+K37MaH3Wmp1TzwY/vATQyWc9qUebGitxWuVD
RdtTEcq6WniDWAJ5FqhHZ3TV/hK7QPTi1gaHG+yJZeXuajsNo6CLrfJy6H6itEfi
XKOns2fiGE/pPxjJpfdTOQipFmw68FuNo8i/A0Nc//d43ejcrqAb9fAKOOTZrpw+
MoqMsFm6V8j+ZN+oKHKSPaD4i6iNuQENBFowLigBCADKSSCJNCY0vPVz8RaCy/uJ
byiZ4dkEUIFkE4TKFCulG8QUMdfczUtYfuUH4ir5vNsG2vxHqDo7W0CBZ1nZjVW9
uUy0TrNrVEsPDcMEqn827oK/pqQmlPq6wxGr6qfrMeAnQKKyQpYA0bwWDxwJ6BBb
0Lw/YyulbLyoCEUPm4Usn+WA8xvUxoWYj/pjg773OLyoznETQiabieNpTmkgad6x
0mH1mbjT0r0RCR0ZUqL1tjGUAfIEr58AVKvP4vZT8jw4quma2QFKLrSswF/bCXqr
K/Eqm+S2lDcOUlY35/fZrBt9Mmr8dF00KYWeND0NE0HFB1cpK5bhHKqMSuwOlrbn
ABEBAAGJATwEGAEIACYWIQTdbDzq2B9JD2WAtKLOaEY1/aXTHwUCWjAuKAIbDAUJ
A8JnAAAKCRDOaEY1/aXTH63LB/4qt+H+3HNEvaRgigod+srkxyT/nQH1tLSHQtht
fukuCgNY7J1y/qGroZxZbB6HSJi//64CH0bV0P06nNoDJt2lPJxKA8nuhxiFEZkf
ACqtJB4W6CUUIZws9YSxVuV84gHZ4g1eQ6mO99R/4jCbhGCebxr0IgPxkulao9Z+
jjb+fdwRkztLKL5GLpiPnR9TuLPxVTB9rnuXsHlGdT4rUXDUKGVdI+wimjjurwvw
vAh3MTVCC0qvQq4V1T1yTCYZ+J7p5wrt1UsBCtYKJfKTeAZN9T7Ji3LVr4jUG2Gn
zHBlhCAdlhsz+4TN+d04QprL2RW86TsIebptwxUscjqJ8lXO
=b72A
-----END PGP PUBLIC KEY BLOCK-----`
}

// private key for testing
func testKeyPrivate() string {
	return `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
=JWlk
-----END PGP PRIVATE KEY BLOCK-----`
}
