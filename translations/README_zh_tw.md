# inject
--
    import "github.com/codegangsta/inject"

inject套件是個提供多種Mapping和dependency inject的工具。

## 使用方法

#### func  InterfaceOf

```go
func InterfaceOf(value interface{}) reflect.Type
```
InterfaceOf提領interface類型的指標。若傳入值非interface的指標則會發生panic。

#### type Applicator

```go
type Applicator interface {
    // Type map對struct中每個用'inject'標記的欄位相依性的對照
    // 注入失敗將回傳error.
    Apply(interface{}) error
}
```

Applicator是用來Mapping相依性到struct的介面。

#### type Injector

```go
type Injector interface {
    Applicator
    Invoker
    TypeMapper
    // SetParent用來設定父injector. 如果在目前injector的Type map中找不到相依，
    // 將會繼續從它的父injector中找，直到回傳error.
    SetParent(Injector)
}
```

Injector是Mapping和inject相依到與函數參數的介面。

#### func  New

```go
func New() Injector
```
New建立並回傳一個Injector.

#### type Invoker

```go
type Invoker interface {
    // Invoke嘗試將interface{}作為一個函數來調用，並基於Type為函數提供參數。
    // 它將回傳reflect.Value的切片，其中存放原函數的回傳值。
    // 如果注入失敗則回傳error.
    Invoke(interface{}) ([]reflect.Value, error)
}
```

Invoker是透過reflection呼叫函數的介面。

#### type TypeMapper

```go
type TypeMapper interface {
    // 使用自身的立即型別從reflect.TypeOf對應interface{}的值。
    Map(interface{}) TypeMapper
    // 使用interface提供的指標對應interface{}的值。
    // 這對於應對一個值為一個interface來說很有用，如果不用指標，
    //interface在這時候無法被直接參照
	MapTo(interface{}, interface{}) TypeMapper
	// 提供使用型別或值直接插入一個mapping的可能方法。
	// 讓直接Mapping可行，用於如單向channels那種無法透過reflect直接對應型別
	// 參數的情境。
	Set(reflect.Type, reflect.Value) TypeMapper
	// 回傳應對到目前型別的值。若型別未對應則回傳零。
    Get(reflect.Type) reflect.Value
}
```

TypeMapper是用於根據型別應對interface{}的值的介面。


## 譯者

Festum Qin (Festum@G.PL)