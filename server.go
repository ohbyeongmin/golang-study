package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	s1 string //test server https address
	s2 string
)

func main() {

	// flagset 만들어주기
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// 처음에는 StringVar 로 하나 받는 것만 해준다. 추후에는 struct 로 만들어 준다.
	// 디폴트 값은 https://daum.net 이다.

	/*
			Command line flag syntax

			The following forms are permitted:

				-flag
				-flag=x
				-flag x  // non-boolean flags only
			One or two minus signs may be used; they are equivalent.
			The last form is not permitted for boolean flags because the
			meaning of the command
				cmd -x *
			where * is a Unix shell wildcard, will change if there is a file
			called 0, false, etc. You must use the -flag=false form to turn
			off a boolean flag.

			Flag parsing stops just before the first non-flag argument
			("-" is a non-flag argument) or after the terminator "--".

			Integer flags accept 1234, 0664, 0x1234 and may be negative.
			Boolean flags may be:
				1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
			Duration flags accept any input valid for time.ParseDuration.

			The default set of command-line flags is controlled by
			top-level functions.  The FlagSet type allows one to define
			independent sets of flags, such as to implement subcommands
			in a command-line interface. The methods of FlagSet are
			analogous to the top-level functions for the command-line
			flag set.

			핵심적인 것은 코드상에서 - 나 -- 를 쓰면 안된다.
			아래 예에서 fs.StringVar(&s1, "-u" 또는 "--u",value,usage) 안된 다는 것이다.
		    하지만 실행단에서 즉 파라미터에서는 "-u" 또는 "--u" 가능하고 또한 동일 하다.
	*/
	fs.StringVar(&s1, "u", "https://daum.net", "https address")

	// 사용자 입력값을 파싱한다. 즉, 입력된 파라미터를 s1 에 집어 넣는다.
	fs.Parse(os.Args[1:]) // command line 의 slice 의 첫번째 파라미터 부터 끝까

	// 여기에 문제가 있다. 찾을 수 있을까?
	if len(os.Args) < 2 {
		// 에러 보고 및 exit
		// 에러 보고 코드 들어가야 함.

		// 사용설명 출력
		fs.Usage()
		os.Exit(1)
	}

	// 여기부터 사용자의 입력 파라미터를 사용하는 코드가 들어간다. 실제적인 코딩 부분
	fmt.Println("사용자 입력 파라미터", s1)
}

// https://pkg.go.dev/flag#NewFlagSet
// https://golang.org/src/flag/flag.go
// 추후 go mod 나 기타 go dep 사용 가능. 아직까지 dependency issue 가 없음.
