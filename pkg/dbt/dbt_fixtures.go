// Copyright © 2019 Nik Ogura <nik.ogura@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbt

//func testDbtConfigContents(port int) string {
//	return fmt.Sprintf(`{
//  "dbt": {
//    "repository": "http://localhost:%d/dbt",
//    "truststore": "http://localhost:%d/dbt/truststore"
//  },
//  "tools": {
//    "repository": "http://localhost:%d/dbt-tools"
//  }
//}`, port, port, port)
//}
//
//func testDbtConfig(port int) Config {
//	return Config{
//		Dbt: DbtConfig{
//			Repo:       fmt.Sprintf("http://localhost:%d/dbt", port),
//			TrustStore: fmt.Sprintf("http://localhost:%d/dbt/truststore", port),
//		},
//		Tools: ToolsConfig{
//			Repo: fmt.Sprintf("http://localhost:%d/dbt-tools", port),
//		},
//		Username: "",
//		Password: "",
//	}
//}

//func testTruststore() string {
//	return fmt.Sprintf("%s\n%s", testKeyPublic(), testkey2Public())
//}
//
//// public key for testing
//func testKeyPublic() string {
//	return `-----BEGIN PGP PUBLIC KEY BLOCK-----
//
//mQENBFowLigBCAC++pVrVRRM86Wo8V7XJsOmU2xtBBY5a8ktB1tdpEhzlPWQHObx
//LINj79HE3lRlIFQmxnKcX3I15bzT3yo3XWLyVUsCDA1Mg9JoU2zJ+u3XftdNBg8J
//eRlTiEwZYflxEYZFSyh3TZI2VZxxlINp/jOGG0dpAdKF3sfKxdTRb30lgDr+wIzv
//oncrjX023UQDHoRZ3f+zPpnkubjhwH8jUHLiGsyKvu0XDB0c4y/6yG6vLUMQDuKX
//bkzBtssdLLA6MTur9Q26dQV/DvuNZdHx17vwXSvf/JMKdWcX80fsAJD644KW9DOg
//pgLqtBa4Tfutt3S8ueIHDnPZBKFL0u+Q61xvABEBAAG0HERCVCBUZXN0IFN1aXRl
//IDxkYnRAZGJ0LmNvbT6JAVQEEwEIAD4WIQTdbDzq2B9JD2WAtKLOaEY1/aXTHwUC
//WjAuKAIbAwUJA8JnAAULCQgHAgYVCAkKCwIEFgIDAQIeAQIXgAAKCRDOaEY1/aXT
//H52LCACYqQnVmJRarckqh1//FUFFpXlTcwWV2zGr3CEFRs0BrWEQD7giehFpKoTL
//JOJJSFd4xcbo/9wMXpJ16soK83o48laxkj+2LDUfDylnTVpVI6zVvAseqnt5nbrA
//CWes75FeIHtQ6woDy7K3RHUORNZ+K37MaH3Wmp1TzwY/vATQyWc9qUebGitxWuVD
//RdtTEcq6WniDWAJ5FqhHZ3TV/hK7QPTi1gaHG+yJZeXuajsNo6CLrfJy6H6itEfi
//XKOns2fiGE/pPxjJpfdTOQipFmw68FuNo8i/A0Nc//d43ejcrqAb9fAKOOTZrpw+
//MoqMsFm6V8j+ZN+oKHKSPaD4i6iNuQENBFowLigBCADKSSCJNCY0vPVz8RaCy/uJ
//byiZ4dkEUIFkE4TKFCulG8QUMdfczUtYfuUH4ir5vNsG2vxHqDo7W0CBZ1nZjVW9
//uUy0TrNrVEsPDcMEqn827oK/pqQmlPq6wxGr6qfrMeAnQKKyQpYA0bwWDxwJ6BBb
//0Lw/YyulbLyoCEUPm4Usn+WA8xvUxoWYj/pjg773OLyoznETQiabieNpTmkgad6x
//0mH1mbjT0r0RCR0ZUqL1tjGUAfIEr58AVKvP4vZT8jw4quma2QFKLrSswF/bCXqr
//K/Eqm+S2lDcOUlY35/fZrBt9Mmr8dF00KYWeND0NE0HFB1cpK5bhHKqMSuwOlrbn
//ABEBAAGJATwEGAEIACYWIQTdbDzq2B9JD2WAtKLOaEY1/aXTHwUCWjAuKAIbDAUJ
//A8JnAAAKCRDOaEY1/aXTH63LB/4qt+H+3HNEvaRgigod+srkxyT/nQH1tLSHQtht
//fukuCgNY7J1y/qGroZxZbB6HSJi//64CH0bV0P06nNoDJt2lPJxKA8nuhxiFEZkf
//ACqtJB4W6CUUIZws9YSxVuV84gHZ4g1eQ6mO99R/4jCbhGCebxr0IgPxkulao9Z+
//jjb+fdwRkztLKL5GLpiPnR9TuLPxVTB9rnuXsHlGdT4rUXDUKGVdI+wimjjurwvw
//vAh3MTVCC0qvQq4V1T1yTCYZ+J7p5wrt1UsBCtYKJfKTeAZN9T7Ji3LVr4jUG2Gn
//zHBlhCAdlhsz+4TN+d04QprL2RW86TsIebptwxUscjqJ8lXO
//=b72A
//-----END PGP PUBLIC KEY BLOCK-----`
//}
//
//// private key for testing
////func testKeyPrivate() string {
////	return `-----BEGIN PGP PRIVATE KEY BLOCK-----
////
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
////=JWlk
////-----END PGP PRIVATE KEY BLOCK-----`
////}
//
//func testkey2Public() string {
//	return `-----BEGIN PGP PUBLIC KEY BLOCK-----
//
//mQENBFpAPGYBCACtRHQZMgHhmETN6X6MCkP7H88jVBSTwhMoZgk0vl6BWK832Uvi
//SMGiZ63uiPkzoUwOtFhexE0QYgKvGPLTm7RWK2aPmsQOk1o+ksFElsRJxT7LzPEM
//g2ci5qAs7q9H8uEntEqfxb9Yn6yiUOLyw6nCrKc9bJN2dCEszpciZoz7AN1ScU+8
//QM3mBw1ToWUB3AMVkd7jJCVloeYprQbqc7pkJBDy9wAISlNRMeLz0PnEuBIrrz8Z
//An1QcqX0PQVWVqNb/duMK5ZszWGK0owfdeSeQiSLK9kvywwo9KZ5qs8XkCUvldZg
//Qv+5mFKK4/+IVReKlnfMvGKRwrGi1oin1mWXABEBAAG0HURCVCBUZXN0IFN1aXRl
//IDxkYnQyQGRidC5jb20+iQFUBBMBCAA+FiEESN604MD5N5baHiTBzGsmLbfIHwUF
//AlpAPGYCGwMFCQPCZwAFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AACgkQzGsmLbfI
//HwX7sQf+Pt6uy3ZZGWZRldGR4qKT/qUCEAc/AG0b2sSlwDt79tEIkprmhMlvNgqF
//DnK5MmAJjZVeR9urWeYeXYk3CH+ZOcR2AOr/15+1LKzkmmVfWGTagQPUFIKkaBdi
//7ymM2wFCvYWG149X1w/ThZ1ZRSCmpuyHiheW/WAE9P1LFTHI+feIjB7Iflmnxkwq
//tPxYxsc07IU0ZIs5+uuerTLlj8gS8acIeGkpNM97m+CiMoL4RLMyt1qFp/lFbbv1
//Td6Dn6vzNL1ZcYGNQ/SHfLTdcnM9GSA5IA1+RyVWb8npFG/sKdPgHMSYOusbdncP
//70qxf+aa/xvZbTvniXtgmhrqog6DJLkBDQRaQDxmAQgA3s4IfL2wUwRN0aCKFeuW
//yOEgNdMuIS9ZKA+f8/s+VXoLtJQ32gZpqCEl2ESOggJx7+ThAhf4F/SDnEdRHZeJ
//IhUnRmzzQPSrhWo9UWDff6cooO1CiSGDSRBgAT8RvZib3OnRWSe67s0webpytUiO
//+y4gt0FtEhXC9kdJ1DsjzJGVyXFYR9pOTV3xjfGKBhHH/6c6kYr1UL2boy8t3IZP
//jyhBHrp3EgxYV7g96ncAVXha91mZ6IisGyXtsOL5qEwPPJCKD2QTKwkJ4S2qqcAR
//8n7agRD8Cn2HESgPezXeg2uoaStcHzhNhF/o/71j2oj5c2u5HkchAzj3l+XHQIrs
//VwARAQABiQE8BBgBCAAmFiEESN604MD5N5baHiTBzGsmLbfIHwUFAlpAPGYCGwwF
//CQPCZwAACgkQzGsmLbfIHwVwjAf9GWR6LdtoEXYVRUSSB2ccl2IeRiwcaEZl/96A
//I2hFX+SCqBVnwJN4jgvhPlCF6PXylkjZKUczrAaizjuU2ZuAt6ONDkEc0R5Glt7j
//dgyl/51WEdBbYeLuVfONtAOBqzs3iRrK8WHnoV+SYQy5aT4kTPTDVzrz/EBDF7KQ
//jqZ4J0i6qsp2DiOxhPn/xVk/iaTRDvtvsA37Qw0mqRlf6xSSLQabroNtJENmf7Cc
//f51a/98jWbflcGLSg/BG2K4hba7ZNKIgKYrS+SKqx5YeE70y/rbjQcJ0ai09Fojc
//hxfIreyexqK3w7pLJFaTbs4ykxWvQZyF0s7h60THq9g76lTjLQ==
//=KIOK
//-----END PGP PUBLIC KEY BLOCK-----
//`
//}
//
////func testkey2Private() string {
////	return `-----BEGIN PGP PRIVATE KEY BLOCK-----
////
////lQOYBFpAPGYBCACtRHQZMgHhmETN6X6MCkP7H88jVBSTwhMoZgk0vl6BWK832Uvi
////SMGiZ63uiPkzoUwOtFhexE0QYgKvGPLTm7RWK2aPmsQOk1o+ksFElsRJxT7LzPEM
////g2ci5qAs7q9H8uEntEqfxb9Yn6yiUOLyw6nCrKc9bJN2dCEszpciZoz7AN1ScU+8
////QM3mBw1ToWUB3AMVkd7jJCVloeYprQbqc7pkJBDy9wAISlNRMeLz0PnEuBIrrz8Z
////An1QcqX0PQVWVqNb/duMK5ZszWGK0owfdeSeQiSLK9kvywwo9KZ5qs8XkCUvldZg
////Qv+5mFKK4/+IVReKlnfMvGKRwrGi1oin1mWXABEBAAEAB/48q4+Lkvsdp1fQUkZB
////ewa61DSPxk0+P+K9bp0intYwrIEOKURtA9TprSw2Ljg0X/Jl47hd1pa1edw+u5mr
////hwXqdl//QM4N3ILpDjImdjRdMHgAnM569zgR/HdxxFIT+3PjlznaIEAUJDUq0O/m
////2aSDyjj1RSONgrRrIBvDq+0JpnwGmuLTXYD39vDtrt2y+Ev/r7O6aJ/AGt4heBqf
/////7Lq4C73hod5YzQ/bjIIAqw2co336UQsxxUeZ4sGiNHr5G8r82tVlnr9e4KOcDQj
////OpCHyDf8iF+sxYUCR58duMv7LUnOix5ILA50SCna+gHos2/XgK9EhQSlPWgeXOJH
////ba3BBADFBfycgc2aurAYAR6ThD9zPGGtSycLwb9rxV2CC6vKKi2tsfvwbgWNekHL
////ZWvYVtCcd5dOf4aINYK7P/0iecrIi3LDu4XuZupxs2TggUUiM9/gQ64D0AgdWTGX
////0vrEV483x4VqXfEc4sjWkyfJ1ME5z2+6aVsK8LMN/1znkNvHiQQA4SIIKXQCb03g
////6+ZkhsWcN2gSv9O8BCsmucv9GZinJyDOKYPxTLyNtmpYAHYzpEAXKB8yRCepAVCz
////Ns6fjU/emHTcC+O5r3my7CklcTaXWXtJl83vfCkhSjNxHVElpfKi2hY3R20qO13a
////nbZGfEOmJ8WWT2FHoG2Wm4UTb7BkXB8D/R4tnaQKkKnmdZrgN4G2xv3jblMG5y5Q
////KrOn2KWkX/+9Mhl8qgFyMcf6npC1Qq/wtK83jHQibsjbMig9AppN+4HEh+sz4vnY
////j+gmRpxVONm4W5DZ/W6hEip1F1ZYYfm9fP3XYLjHLHEdrcEOF5ps8i7mdx2xePWV
////GF/yMXY3WTxpQcO0HURCVCBUZXN0IFN1aXRlIDxkYnQyQGRidC5jb20+iQFUBBMB
////CAA+FiEESN604MD5N5baHiTBzGsmLbfIHwUFAlpAPGYCGwMFCQPCZwAFCwkIBwIG
////FQgJCgsCBBYCAwECHgECF4AACgkQzGsmLbfIHwX7sQf+Pt6uy3ZZGWZRldGR4qKT
/////qUCEAc/AG0b2sSlwDt79tEIkprmhMlvNgqFDnK5MmAJjZVeR9urWeYeXYk3CH+Z
////OcR2AOr/15+1LKzkmmVfWGTagQPUFIKkaBdi7ymM2wFCvYWG149X1w/ThZ1ZRSCm
////puyHiheW/WAE9P1LFTHI+feIjB7IflmnxkwqtPxYxsc07IU0ZIs5+uuerTLlj8gS
////8acIeGkpNM97m+CiMoL4RLMyt1qFp/lFbbv1Td6Dn6vzNL1ZcYGNQ/SHfLTdcnM9
////GSA5IA1+RyVWb8npFG/sKdPgHMSYOusbdncP70qxf+aa/xvZbTvniXtgmhrqog6D
////JJ0DmARaQDxmAQgA3s4IfL2wUwRN0aCKFeuWyOEgNdMuIS9ZKA+f8/s+VXoLtJQ3
////2gZpqCEl2ESOggJx7+ThAhf4F/SDnEdRHZeJIhUnRmzzQPSrhWo9UWDff6cooO1C
////iSGDSRBgAT8RvZib3OnRWSe67s0webpytUiO+y4gt0FtEhXC9kdJ1DsjzJGVyXFY
////R9pOTV3xjfGKBhHH/6c6kYr1UL2boy8t3IZPjyhBHrp3EgxYV7g96ncAVXha91mZ
////6IisGyXtsOL5qEwPPJCKD2QTKwkJ4S2qqcAR8n7agRD8Cn2HESgPezXeg2uoaStc
////HzhNhF/o/71j2oj5c2u5HkchAzj3l+XHQIrsVwARAQABAAf/bx0KH4eMBUK6umXZ
////J4hnmMOpCB7KvRL7YB62svrjwcYNS1+1oN4c2BO752h9N3oXrz8SNbSVbgDrl8b7
////PYpCggRRo24XPBgo5+6tKMXqgCNxgBwC2Bel/QyVDFoTM14FsPzMgi1trMUYvURO
////C3llkP+Woj2XMvM2WRMBTz21I0GzXMYsns1B6E5vsH00hWSGm52wq6USt4k65Lu3
////oks7KDC8yOWRQ+1GFwMZ+NPYors1YadSTay44tp9xC7Xur8zHpCiIMCmYDVdxZuM
////yQQblB2pkhwuxndpkTPcHXwj0oHVaje5+bwWDeiB9EhYQoLt285PaWAowF2q4GLH
////roQ+oQQA7LV7oR0spfdphanApJMsQsbdrd3jPCgzku6S9vbQccYz7MpxZPecPBnU
////4VYHjNjQA0WvKPrpMCYVuHWECtGNlgBmWj3+kQL3plYHswxeG2+yLpcOTHqNWb7n
////fw8wd6k12izay+UK+5rEr+nNzbf6aoKa1chk8dHUyxODuR8JVncEAPD2dyvWv36x
////Vg9bQLTdoOLpWoPUIKxH0co3blrjcqSBTDxpOjD00TJM4N/4UMBLNQJUlpbotC+A
////cHaJIgzSW7hPRkv9OOJw6f6Ti94ZaA1XbI4Wxrs2VKWUaWK8K82WvPOb4l8Y0i0q
////W9iFm7WtqwP8nVbeyGrtxrMTGvYQBjEhA/4szhAw/kUh2BzIE3DHhQtVZi3qURXw
////KBTh1h6KfVVXGol4WCqJ+xHT6LFadvdz4Dn+4G6pltW/bQ76d0yGcCUJucmpEnS9
////82fgUyM+n2goVn+6v+PE8xwLcXmMiqlKyueKZM2weElXHJC954wObanQAnEufYU7
////MxSCOII6nDkIwEFJiQE8BBgBCAAmFiEESN604MD5N5baHiTBzGsmLbfIHwUFAlpA
////PGYCGwwFCQPCZwAACgkQzGsmLbfIHwVwjAf9GWR6LdtoEXYVRUSSB2ccl2IeRiwc
////aEZl/96AI2hFX+SCqBVnwJN4jgvhPlCF6PXylkjZKUczrAaizjuU2ZuAt6ONDkEc
////0R5Glt7jdgyl/51WEdBbYeLuVfONtAOBqzs3iRrK8WHnoV+SYQy5aT4kTPTDVzrz
/////EBDF7KQjqZ4J0i6qsp2DiOxhPn/xVk/iaTRDvtvsA37Qw0mqRlf6xSSLQabroNt
////JENmf7Ccf51a/98jWbflcGLSg/BG2K4hba7ZNKIgKYrS+SKqx5YeE70y/rbjQcJ0
////ai09FojchxfIreyexqK3w7pLJFaTbs4ykxWvQZyF0s7h60THq9g76lTjLQ==
////=Gjbj
////-----END PGP PRIVATE KEY BLOCK-----
////`
////}
