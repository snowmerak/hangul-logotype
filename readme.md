# Hangul Logotype

`hangul-logotype`은 한글 자모(자음, 모음) 입력을 실시간으로 완성형 한글로 조합해주는 Go 라이브러리입니다. 영어 QWERTY 키보드 입력을 한글로 자동 변환하는 기능도 포함하고 있어, 한/영 전환 없이도 한글 입력이 가능합니다.

## 주요 기능

- **실시간 한글 조합**: 자모 입력을 받아 완전한 한글 음절로 조합합니다.
- **영어 키보드 매핑**: 기본 두벌식(QWERTY) 레이아웃에 맞춰 영어 알파벳 입력을 한글 자모로 자동 변환합니다.
  - `gksrmf` → `한글`
  - 대문자 입력 시 쌍자음/이중모음으로 변환 (`R` → `ㄲ`, `O` → `ㅒ`)
- **다중 키보드 레이아웃 지원**: `KeyboardLayout`을 통해 키 매핑을 교체할 수 있으며, 두벌식 외에 **세벌식 최종** 레이아웃(`SebulshikFinalLayout`)이 기본 제공됩니다.
  - `mfskgw` (세벌식 최종) → `한글`
- **복잡한 한글 규칙 처리**:
  - **이중모음**: `ㅗ` + `ㅏ` → `ㅘ`
  - **겹받침**: `ㄷ`+`ㅏ`+`ㄹ`+`ㄱ` → `닭`
  - **연음 법칙**: 종성이 다음 모음 앞에서 초성으로 이동하는 규칙을 처리합니다. (`ㄱ`+`ㅏ`+`ㄴ`+`ㅏ` → `가나`)
- **스트림 기반 처리**: `io.RuneWriter`와 유사한 인터페이스를 제공하여 문자열 스트림을 쉽게 처리할 수 있습니다.

## 설치

```bash
go get github.com/snowmerak/hangul-logotype/hangul
```

## 사용법

`LogoTyper`를 생성하고 `WriteRune` 메서드로 문자를 입력하면, 내부 버퍼에서 자동으로 한글을 조합합니다. `Result` 메서드로 최종 결과를 얻을 수 있습니다.

```go
package main

import (
	"fmt"
	"github.com/snowmerak/hangul-logotype/hangul"
)

func main() {
	typer := hangul.NewLogoTyper()

	// 영어 문장 입력
	input := "dkssudgktpdy, gksrmf!"

	typer.WriteRunes([]rune(input))

	// 결과 출력
	result := string(typer.Result())
	fmt.Println(result) // 출력: 안녕하세요, 한글!

	// 세벌식 최종 레이아웃 사용
	sebul := hangul.NewLogoTyperWithLayout(hangul.SebulshikFinalLayout)
	sebul.WriteString("mfskgw")
	fmt.Println(string(sebul.Result())) // 출력: 한글
}
```

## API

- `hangul.NewLogoTyper() *LogoTyper`: 새로운 `LogoTyper` 인스턴스를 생성합니다.
- `hangul.NewLogoTyperWithLayout(layout hangul.KeyboardLayout) *LogoTyper`: 지정한 키보드 레이아웃으로 `LogoTyper`를 생성합니다.
- `(lt *LogoTyper) WriteRune(r rune)`: 한 문자를 `LogoTyper`에 씁니다. 영어인 경우 자동으로 한글 자모로 변환됩니다.
- `(lt *LogoTyper) WriteRunes(runes []rune)`: 여러 문자를 한 번에 `LogoTyper`에 씁니다.
- `(lt *LogoTyper) Result() []byte`: 현재까지 조합된 최종 결과물을 `[]byte`로 반환합니다.
- `(lt *LogoTyper) Reset()`: `LogoTyper`의 상태(버퍼, 결과)를 초기화합니다.
- `(lt *LogoTyper) WithLayout(layout hangul.KeyboardLayout) *LogoTyper`: 기존 인스턴스에 새로운 키보드 레이아웃을 적용합니다.
- `hangul.LogoTypeWithLayout(writer *bytes.Buffer, input []rune, layout hangul.KeyboardLayout)`: 주어진 레이아웃을 사용해 자모를 완성형으로 조합합니다.
- `hangul.KeyboardLayout`: 입력 rune을 한글 자모 또는 문자로 매핑하는 타입입니다.
- `hangul.DubeolsikLayout`, `hangul.SebulshikFinalLayout`: 기본 제공 레이아웃 매핑.

## 📄 라이선스

이 프로젝트는 [LICENSE](./LICENSE) 파일에 명시된 라이선스를 따릅니다.
