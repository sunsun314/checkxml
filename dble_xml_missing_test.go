package checkxml

import (
	"fmt"
	"testing"
)

func TestMissingConfig(t *testing.T) {

	data := []byte(`
        <?xml version="1.0" encoding="UTF-8"?>
       <!DOCTYPE dble:rule SYSTEM "rule.dtd">
          <dble:rule>
            <tableRule name="sharding-by-enum">
              <rule>
                <columns>id</columns>
                <algorithm>enum</algorithm>
              </rule>
            </tableRule>

            <function name="hashmod" class="Hash">
              <property name="partitionCount">4</property>
              <property name="partitionLength">1</property>
             </function>
            </dble:rule>
        `)
	tv := CoreRuleXml{}
	mems, _, _ := MissingXMLTags(data, tv)

	fmt.Print(mems)

}
