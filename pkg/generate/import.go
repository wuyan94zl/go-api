package generate

import "fmt"

func getImportStr(name string) string {
	imp := `import (
	"fmt"
	"github.com/gin-gonic/gin"
	"%s"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)
`
	imp = fmt.Sprintf(imp, name)
	return imp
}
