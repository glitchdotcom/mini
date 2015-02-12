package mini

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strings"
	"testing"
)

func TestSimpleIniFile(t *testing.T) {

	simpleIni := `first=alpha
second=beta
third="gamma bamma"
fourth = 'delta'
int=32
float=3.14
true=true
false=false
#comment
; comment`

	filepath := path.Join(os.TempDir(), "simpleini.txt")
	f, err := os.Create(filepath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filepath)
	if _, err := f.WriteString(simpleIni); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	config, err := LoadConfiguration(filepath)

	assert.Nil(t, err, "Simple configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.String("second", ""), "beta", "Read value of second wrong")
	assert.Equal(t, config.String("third", ""), "gamma bamma", "Read value of third wrong")
	assert.Equal(t, config.String("fourth", ""), "delta", "Read value of fourth wrong")
	assert.Equal(t, config.Integer("int", 0), 32, "Read value of int wrong")
	assert.Equal(t, config.Float("float", 0), 3.14, "Read value of float wrong")
	assert.Equal(t, config.Boolean("true", false), true, "Read true wrong")
	assert.Equal(t, config.Boolean("false", true), false, "Read false wrong")

	assert.Equal(t, len(config.Keys()), 8, "Simple ini contains 8 fields")
}

func TestSimpleIniFileFromReader(t *testing.T) {

	simpleIni := `first=alpha
second=beta
third="gamma bamma"
fourth = 'delta'
int=32
float=3.14
true=true
false=false
#comment
; comment`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.String("second", ""), "beta", "Read value of second wrong")
	assert.Equal(t, config.String("third", ""), "gamma bamma", "Read value of third wrong")
	assert.Equal(t, config.String("fourth", ""), "delta", "Read value of fourth wrong")
	assert.Equal(t, config.Integer("int", 0), 32, "Read value of int wrong")
	assert.Equal(t, config.Float("float", 0), 3.14, "Read value of float wrong")
	assert.Equal(t, config.Boolean("true", false), true, "Read true wrong")
	assert.Equal(t, config.Boolean("false", true), false, "Read false wrong")

	assert.Equal(t, len(config.Keys()), 8, "Simple ini contains 8 fields")
}

func TestCaseInsensitive(t *testing.T) {

	simpleIni := `fIrst=alpha
SECOND=beta
Third="gamma bamma"
FourTh = 'delta'`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Case insensitive configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.String("second", ""), "beta", "Read value of second wrong")
	assert.Equal(t, config.String("THIRD", ""), "gamma bamma", "Read value of third wrong")
	assert.Equal(t, config.String("fourth", ""), "delta", "Read value of fourth wrong")

	assert.Equal(t, len(config.Keys()), 4, "Case ins ini contains 4 fields")
}

func TestArrayOfStrings(t *testing.T) {

	simpleIni := `key[]=one
key[]=two
noarray=three`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	val := config.Strings("key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], "one", "Read value of first wrong")
	assert.Equal(t, val[1], "two", "Read value of second wrong")

	val = config.Strings("noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], "three", "Read value of noarray wrong")

	assert.Equal(t, len(config.Keys()), 2, "StringArray test contains 2 fields")
}

func TestArrayOfIntegers(t *testing.T) {

	simpleIni := `key[]=1
key[]=2
noarray=3`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	val := config.Integers("key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], 1, "Read value of first wrong")
	assert.Equal(t, val[1], 2, "Read value of second wrong")

	val = config.Integers("noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], 3, "Read value of noarray wrong")

	assert.Equal(t, len(config.Keys()), 2, "IntArray test contains 2 fields")
}

func TestArrayOfFloats(t *testing.T) {

	simpleIni := `key[]=1.1
key[]=2.2
noarray=3.3`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	val := config.Floats("key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], 1.1, "Read value of first wrong")
	assert.Equal(t, val[1], 2.2, "Read value of second wrong")

	val = config.Floats("noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], 3.3, "Read value of noarray wrong")

	assert.Equal(t, len(config.Keys()), 2, "FloatArray test contains 2 fields")
}

func TestSectionedIniFile(t *testing.T) {

	simpleIni := `first=alpha
second=beta
third="gamma bamma"
fourth = 'delta'
int=32
float=3.14
true=true
false=false

[section]
first=raz
second=dba
int=124
float=1222.7
true=false
false=true
#comment
; comment`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Sectioned configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.String("second", ""), "beta", "Read value of second wrong")
	assert.Equal(t, config.String("third", ""), "gamma bamma", "Read value of third wrong")
	assert.Equal(t, config.String("fourth", ""), "delta", "Read value of fourth wrong")
	assert.Equal(t, config.Integer("int", 0), 32, "Read value of int wrong")
	assert.Equal(t, config.Float("float", 0), 3.14, "Read value of float wrong")
	assert.Equal(t, config.Boolean("true", false), true, "Read true wrong")
	assert.Equal(t, config.Boolean("false", true), false, "Read false wrong")

	assert.Equal(t, config.StringFromSection("section", "first", ""), "raz", "Read value of first from section wrong")
	assert.Equal(t, config.StringFromSection("section", "second", ""), "dba", "Read value of second from section wrong")
	assert.Equal(t, config.IntegerFromSection("section", "int", 0), 124, "Read value of int in section wrong")
	assert.Equal(t, config.FloatFromSection("section", "float", 0), 1222.7, "Read value of float in section wrong")
	assert.Equal(t, config.BooleanFromSection("section", "true", true), false, "Read true in section wrong")
	assert.Equal(t, config.BooleanFromSection("section", "false", false), true, "Read false in section wrong")

	assert.Equal(t, len(config.Keys()), 8, "Section ini contains 6 fields")
	assert.Equal(t, len(config.KeysForSection("section")), 6, "Section in ini contains 4 fields")
}

func TestArrayOfStringsInSection(t *testing.T) {

	simpleIni := `key=nope
noarray=nope
[section]
key[]=one
key[]=two
noarray=three`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	val := config.StringsFromSection("section", "key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], "one", "Read value of first wrong")
	assert.Equal(t, val[1], "two", "Read value of second wrong")

	val = config.StringsFromSection("section", "noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], "three", "Read value of noarray wrong")

	assert.Equal(t, len(config.Keys()), 2, "StringArray section test contains 2 fields")
}

func TestArrayOfIntegersInSection(t *testing.T) {

	simpleIni := `key=nope
noarray=nope
[section]
key[]=1
key[]=2
noarray=3`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	val := config.IntegersFromSection("section", "key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], 1, "Read value of first wrong")
	assert.Equal(t, val[1], 2, "Read value of second wrong")

	val = config.IntegersFromSection("section", "noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], 3, "Read value of noarray wrong")

	assert.Equal(t, len(config.KeysForSection("section")), 2, "IntArray section test contains 2 fields")
}

func TestArrayOfFloatsInSection(t *testing.T) {

	simpleIni := `key=nope
noarray=nope
[section]
key[]=1.1
key[]=2.2
noarray=3.3`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	val := config.FloatsFromSection("section", "key")

	assert.Equal(t, len(val), 2, "Array for keys should have 2 values")
	assert.Equal(t, val[0], 1.1, "Read value of first wrong")
	assert.Equal(t, val[1], 2.2, "Read value of second wrong")

	val = config.FloatsFromSection("section", "noarray")
	assert.Equal(t, len(val), 1, "Array for noarray should have 1 value")
	assert.Equal(t, val[0], 3.3, "Read value of noarray wrong")

	assert.Equal(t, len(config.Keys()), 2, "FloatArray section test contains 2 fields")
}

func TestMultipleSections(t *testing.T) {

	simpleIni := `first=alpha
int=32
float=3.14

[section_one]
first=raz
int=124
float=1222.7

[section_two]
first=one
int=555
float=124.3`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Sectioned configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.Integer("int", 0), 32, "Read value of int wrong")
	assert.Equal(t, config.Float("float", 0), 3.14, "Read value of float wrong")

	assert.Equal(t, config.StringFromSection("section_one", "first", ""), "raz", "Read value of first from section wrong")
	assert.Equal(t, config.IntegerFromSection("section_one", "int", 0), 124, "Read value of int in section wrong")
	assert.Equal(t, config.FloatFromSection("section_one", "float", 0), 1222.7, "Read value of float in section wrong")

	assert.Equal(t, config.StringFromSection("section_two", "first", ""), "one", "Read value of first from section wrong")
	assert.Equal(t, config.IntegerFromSection("section_two", "int", 0), 555, "Read value of int in section wrong")
	assert.Equal(t, config.FloatFromSection("section_two", "float", 0), 124.3, "Read value of float in section wrong")

	assert.Equal(t, len(config.Keys()), 3, "Section ini contains 3 fields")
	assert.Equal(t, len(config.KeysForSection("section_one")), 3, "Section in ini contains 3 fields")
	assert.Equal(t, len(config.KeysForSection("section_two")), 3, "Section in ini contains 3 fields")
}

func TestSplitSection(t *testing.T) {

	simpleIni := `first=alpha
int=32
float=3.14

[section_one]
first=raz

[section_two]
first=one
int=555
float=124.3

[section_one]
int=124
float=1222.7`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Sectioned configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "alpha", "Read value of first wrong")
	assert.Equal(t, config.Integer("int", 0), 32, "Read value of int wrong")
	assert.Equal(t, config.Float("float", 0), 3.14, "Read value of float wrong")

	assert.Equal(t, config.StringFromSection("section_one", "first", ""), "raz", "Read value of first from section wrong")
	assert.Equal(t, config.IntegerFromSection("section_one", "int", 0), 124, "Read value of int in section wrong")
	assert.Equal(t, config.FloatFromSection("section_one", "float", 0), 1222.7, "Read value of float in section wrong")

	assert.Equal(t, config.StringFromSection("section_two", "first", ""), "one", "Read value of first from section wrong")
	assert.Equal(t, config.IntegerFromSection("section_two", "int", 0), 555, "Read value of int in section wrong")
	assert.Equal(t, config.FloatFromSection("section_two", "float", 0), 124.3, "Read value of float in section wrong")

	assert.Equal(t, len(config.Keys()), 3, "Section ini contains 3 fields")
	assert.Equal(t, len(config.KeysForSection("section_one")), 3, "Section in ini contains 3 fields")
	assert.Equal(t, len(config.KeysForSection("section_two")), 3, "Section in ini contains 3 fields")
}

func TestRepeatedKey(t *testing.T) {

	simpleIni := `first=alpha
first=beta`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "beta", "Read value of first wrong")

	assert.Equal(t, len(config.Keys()), 1, "ini contains 1 fields")
}

func TestDefaults(t *testing.T) {

	simpleIni := `first=alpha
third=\`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Equal(t, config.String("second", "beta"), "beta", "Read default value of first wrong")
	assert.Equal(t, config.String("third", "gamma"), "gamma", "Read default value of too short a string")
	assert.Equal(t, config.Integer("int", 32), 32, "Default value of int wrong")
	assert.Equal(t, config.Float("float", 3.14), 3.14, "Default value of float wrong")
	assert.Equal(t, config.Boolean("bool", true), true, "Default value of bool wrong")

	assert.Equal(t, config.String("", "test"), "test", "Nil key should result in empty value")
	assert.Equal(t, config.Integer("", 32), 32, "Default value of int wrong for empty key")
	assert.Equal(t, config.Float("", 3.14), 3.14, "Default value of float wrong for empty key")
	assert.Equal(t, config.Boolean("", true), true, "Default value of bool wrong for empty key")

	assert.Equal(t, len(config.Keys()), 2, "ini contains 2 fields")
}

func TestDefaultsOnParseError(t *testing.T) {

	simpleIni := `first=alpha
int=yex
float=blap
bool=zipzap
intarray[]=blip
floatarray[]=blap`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Equal(t, config.String("second", "beta"), "beta", "Read default value of first wrong")
	assert.Equal(t, config.Integer("int", 32), 32, "Default value of int wrong")
	assert.Equal(t, config.Float("float", 3.14), 3.14, "Default value of float wrong")
	assert.Equal(t, config.Boolean("bool", true), true, "Default value of bool wrong")

	assert.Nil(t, config.Integers("intarray"), "Default value of ints wrong on parse error")
	assert.Nil(t, config.Floats("floatarray"), "Default value of floats wrong on parse error")

	assert.Equal(t, len(config.Keys()), 6, "ini contains 4 fields")
}

func TestMissingArray(t *testing.T) {

	simpleIni := `first=alpha`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Nil(t, config.Strings("second"), "Read default value of strings wrong")
	assert.Nil(t, config.Integers("int"), "Default value of ints wrong")
	assert.Nil(t, config.Floats("float"), "Default value of floats wrong")

	assert.Nil(t, config.Strings(""), "Read default value of strings wrong for empty key")
	assert.Nil(t, config.Integers(""), "Default value of ints wrong for empty key")
	assert.Nil(t, config.Floats(""), "Default value of floats wrong for empty key")

	assert.Equal(t, len(config.Keys()), 1, "ini contains 1 fields")
}

func TestDefaultsWithSection(t *testing.T) {

	simpleIni := `[section]
first=alpha`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Equal(t, config.StringFromSection("section", "second", "beta"), "beta", "Read default value of first wrong")
	assert.Equal(t, config.IntegerFromSection("section", "int", 32), 32, "Default value of int wrong")
	assert.Equal(t, config.FloatFromSection("section", "float", 3.14), 3.14, "Default value of float wrong")
	assert.Equal(t, config.BooleanFromSection("section", "bool", true), true, "Default value of bool wrong")

	assert.Equal(t, config.StringFromSection("section-1", "second", "beta"), "beta", "Missing section for first wrong")
	assert.Equal(t, config.IntegerFromSection("section-1", "int", 32), 32, "Missing section for int wrong")
	assert.Equal(t, config.FloatFromSection("section-1", "float", 3.14), 3.14, "Missing section for float wrong")
	assert.Equal(t, config.BooleanFromSection("section-1", "bool", true), true, "Missing section for bool wrong")

	assert.Equal(t, len(config.Keys()), 0, "ini contains 0 fields")
	assert.Equal(t, len(config.KeysForSection("section")), 1, "section contains 1 field")
	assert.Nil(t, config.KeysForSection("section-1"), "missing section should have no keys")
}

type testStruct struct {
	First        string
	Second       string
	L            int64
	F64          float64
	Flag         bool
	Strings      []string
	LS           []int64
	F64s         []float64
	Missing      string
	MissingInt   int64
	MissingArray []string
	Flags        []bool
	U            uint64
	private      string
}

func TestLoadStructFile(t *testing.T) {

	simpleIni := `[section]
first=alpha
second=beta
third="gamma bamma"
fourth = 'delta'
l=-32
f64=3.14
flag=true
unflag=false
strings[]=one
strings[]=two
LS[]=1
LS[]=2
F64s[]=11.0
F64s[]=22.0
#comment
; comment`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	var data testStruct

	data.MissingInt = 33
	data.Missing = "hello world"

	ok := config.DataFromSection("section", &data)

	assert.Equal(t, ok, true, "load should succeed")

	assert.Equal(t, data.First, "alpha", "Read value of first wrong")
	assert.Equal(t, data.Second, "beta", "Read value of second wrong")
	assert.Equal(t, data.L, -32, "Read value of int wrong")
	assert.Equal(t, data.F64, 3.14, "Read value of float wrong")
	assert.Equal(t, data.Flag, true, "Read true wrong")
	assert.Equal(t, data.Missing, "hello world", "Read value of missing wrong")
	assert.Equal(t, data.MissingInt, 33, "Read false wrong")
	assert.Nil(t, data.MissingArray, "Missing array Should be nil")
	assert.Equal(t, data.private, "", "private value in struct should be ignored")

	strings := data.Strings
	assert.NotNil(t, strings, "strings should not be nil")
	assert.Equal(t, len(strings), 2, "Read wrong length of string array")
	assert.Equal(t, strings[0], "one", "Read string array wrong")
	assert.Equal(t, strings[1], "two", "Read string array wrong")

	assert.NotNil(t, data.LS, "ints should not be nil")
	assert.Equal(t, len(data.LS), 2, "Read wrong length of ints array")
	assert.Equal(t, data.LS[0], 1, "Read ints array wrong")
	assert.Equal(t, data.LS[1], 2, "Read ints array wrong")

	assert.NotNil(t, data.F64s, "floats should not be nil")
	assert.Equal(t, len(data.F64s), 2, "Read wrong length of floats array")
	assert.Equal(t, data.F64s[0], 11.0, "Read floats array wrong")
	assert.Equal(t, data.F64s[1], 22.0, "Read floats array wrong")
}

func TestLoadStructMissingSection(t *testing.T) {

	simpleIni := ``

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	var data testStruct

	ok := config.DataFromSection("section", &data)

	assert.Equal(t, ok, false, "section is missing so ok should be false")
}

func TestMissingSection(t *testing.T) {
	simpleIni := `[section]
first=alpha`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	var data testStruct

	val := config.DataFromSection("missing_section", &data)
	assert.False(t, val)
}

func TestMissingArrayInSection(t *testing.T) {

	simpleIni := `[section]
first=alpha`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Nil(t, config.StringsFromSection("section", "second"), "Read default value of strings wrong")
	assert.Nil(t, config.IntegersFromSection("section", "int"), "Default value of ints wrong")
	assert.Nil(t, config.FloatsFromSection("section", "float"), "Default value of floats wrong")

	assert.Nil(t, config.StringsFromSection("section-1", "second"), "Missing section for strings wrong")
	assert.Nil(t, config.IntegersFromSection("section-1", "int"), "Missing section for ints wrong")
	assert.Nil(t, config.FloatsFromSection("section-1", "float"), "Missing section for floats wrong")

	assert.Equal(t, len(config.Keys()), 0, "ini contains 0 fields")
	assert.Equal(t, len(config.KeysForSection("section")), 1, "section contains 1 field")
	assert.Nil(t, config.KeysForSection("section-1"), "missing section should have no keys")
}

func TestBadSection(t *testing.T) {

	simpleIni := `key=nope
noarray=nope
[section
key[]=one
key[]=two
noarray=three`

	_, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.NotNil(t, err, "Configuration should load with error.")
}

func TestBadKeyValue(t *testing.T) {

	simpleIni := `key=nope
noarray:nope`

	_, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.NotNil(t, err, "Configuration should load with error.")
}

func TestBadArrayAsSingle(t *testing.T) {

	simpleIni := `key=nope
noarray=nope
[section]
key[]=one
key[]=two
noarray=three`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Configuration should load without error.")

	assert.Equal(t, config.StringFromSection("section", "key", "default"), "default", "Read default value of strings wrong")
}

func TestBadFile(t *testing.T) {
	const filepath = "/no.such.dir/xxx.no-such-file.txt"
	config, err := LoadConfiguration(filepath)

	assert.NotNil(t, err, "No valid file is an error.")
	assert.True(t, os.IsNotExist(err), "No valid file errors is of expected type.")
	assert.Nil(t, config, "Configuration should be nil.")

}

func TestNullValuesInGet(t *testing.T) {

	assert.Nil(t, get(nil, "foo"), "Configuration should be nil.")

}

func TestStringEscape(t *testing.T) {

	simpleIni := `first=\n\t\rhello`

	config, err := LoadConfigurationFromReader(strings.NewReader(simpleIni))

	assert.Nil(t, err, "Simple configuration should load without error.")

	assert.Equal(t, config.String("first", ""), "\n\t\rhello", "Read value of first wrong")

	assert.Equal(t, len(config.Keys()), 1, "ini contains 1 fields")
}

func BenchmarkLoadConfiguration(b *testing.B) {

	simpleIni := `first=alpha
int=32
float=3.14

[section_one]
first=raz

[section_two]
first=one
int=555
float=124.3

[section_one]
int=124
float=1222.7`

	b.ReportAllocs()
	r := strings.NewReader(simpleIni)
	for i := 0; i < b.N; i++ {
		r.Seek(0, os.SEEK_SET)
		_, err := LoadConfigurationFromReader(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}
