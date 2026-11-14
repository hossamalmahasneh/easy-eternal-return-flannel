package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.0" )

func np8iy45fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxaMS8THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKNgUwTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EW6A4HJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xo358CfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MkisvvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUlwTJm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVn0JAmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHrY9OIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkQYhDRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f81BEBb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kqM9XMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjXD1io7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWcNzYSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8d5YWp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kjZYrd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTXZQAFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vhlvXjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTHE4YA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6Ro0tA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9169OhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zw5SCeFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8x2QqiFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqSOObFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JihxAgM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fO0XAMSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCVosswrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLo35j5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8m5aUvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tvv66aR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYMvCz7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKMRc0cwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h92aawLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrPoDI64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYeUmDhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymhCwOGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JV7jzPWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6NqwFnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPiLCk7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XsLuk77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vh8CPNZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxMfwkayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVBmUdMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmLzkrzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8lZWQfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4oNsPcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mysvm7stWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOnMtRHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53SwtN8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmW9rnmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dNaFW7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FY5BuevtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ryg7YXHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6kVGDwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEYP0SX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqywKBgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mw1ljJ4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSdTTXbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQNoNfNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPhsSR4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcj9xolRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoQXrpIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4iDqVnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBG3z46MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEVv5hP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUextettWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmlkVezoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peeF3hlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnb65angWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CL6B46QZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6fm4yPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfwiPvfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANXvheqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kb7DMDFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOGKX9V4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Veke7qwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlgfZN81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi2DDCIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjJJE0mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2ZI3bO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF2dIAWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taMx3YSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2rGoLJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL6JwyPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PTpMz7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kSX27fbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzgvGZj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQQQ51CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trlJKwfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLjW5zGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRd2lrvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMqsdy7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5sniJkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9L3ZMw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqgML6WmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4i4t0mLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voSlypOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYWAaeUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kY4iENTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyLIrrudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1u7x4dJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LPy4xrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbZjqLtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eKyTjA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuM0q7BWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vl68ckRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xmrths93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kifPh90fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASb10DjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y57olSSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FR0wgcBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZG4rvTwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAF7mZbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grzaGJFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oV8SYemEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57DxFx26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xlZbLpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f51qX44yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qRHvWMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4vreMotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcbJpP8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gORTHWJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QB7sPRreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOTjNozhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAwh2LgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVjhD5j2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PPeGmCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqRtdLzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JR4lmjpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F36MNeNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxqfY3MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsgLzqmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5eWoKooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJrhCt6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivT2elZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxUVg9mFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qmvlo94LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqOjlWQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHfsJzZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2p78A6iJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V83uVwYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6d9M5qnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4jMcOrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaWAGP5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgB6lpJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64cZOnE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDNHGreRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQv0Qz9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWTRE8c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFbEiS2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzyKadTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dQg1gREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dg81xcTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vt99H1OMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6SUz13iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4VF9eshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qflnc7KoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Airws7JNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVaI6dYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3OlNmTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwhHwy3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJ6Sk0XoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJxcQBxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP2TyMLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvmFYebgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L03sJHOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkgeRuWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKIeXRGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TGiNkYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwUM4NZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOcawI7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tc2AAcJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHZ4LFb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFDLTK6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtxXNiJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoXHVBnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZX78S4cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkOWNbU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eYIpXPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EebhXTdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Klx9kDVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2e49pLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qMJTGLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8NADo4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv5dRtbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDymu2hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qA6CJblzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdjA2VrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNVII58RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cw7TOTC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCP6ocDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hmwzLWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tF5Q2FZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHnK1Q9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ao2tAHBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNisoFh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veJUEVgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEGskrqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuoc9UzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZ0EsVcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DDSMPDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXd4O3osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otvC9MfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kifGDNYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6ozG1PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trv4TfjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUkZPi6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2TETaMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0PECHnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dHFUolpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOGXD7ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFfhfUkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poWgQXmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQud2YGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gc8hiY4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ATBf8nsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D09EqMVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBtBx12zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNEH3A2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKLp35RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kx8wKWxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZvgzzxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKPvIJw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7LRDvMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OinK01DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wa2wPajXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Om32vu6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3lNCMXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02XJY3sRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJnBibBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eFuABWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltnyF1wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uc6FOdSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNChCNOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jyqx4JvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu64xvboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ekVRf4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlWkUXKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neU5WxijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAAq6RB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZPuyLdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuZLfDwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCpKKgF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvD1gxIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRGLytPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FfYWlyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Juk4s0qXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yuj4MNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byX7bk5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWrBC9k3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHL1Wy74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Qu9lAyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wYlpxVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjpYmblyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkrPdubKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzJoKD4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cazmFKiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 976DgwKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfDrLGcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgV7R3sKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0NBPNrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhwlGnmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLWRdfdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkczEeF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eIKHTUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXQi2LiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PGGbZ6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlcvxuBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaLg29HlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLp32v2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l2lggYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ui35Ffz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LnNrTyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AntA5QMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEso1HCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcGu3bqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeOtWHXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N2WzW1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGc7pAQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bh0JfwSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8Ad89DxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUy38sMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6MinBVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfKenIgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdn72VxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6G9SYILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dy8Ow2NsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsInSzufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0scQVXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lz7faqAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNA258AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbOAKCwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDCODpAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pqOHVRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwPxpKe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxgH83lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KxQHdHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvMh5xkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDav9QlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OC9rYJ1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSdDtKAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGhsdH6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jkp5ZUbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoiEivJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0Ue3cbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O11WwTrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FBgYJ6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5DFetKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rd8iJIABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtuGtOeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JszKWr1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLkQJcCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyEd5dXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CX1YRWZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62cQHAzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dqrok0ZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI3rWeWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5BpUAHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoI4kNZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAKIFjKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbxy2yBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GRpABynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WE2CjmlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rb3nERwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptUkr27dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozWFSakDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCFSo3jdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZajLQS2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wrzrSobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeZ50GUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNyErR97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCQJVrnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5rmT2xDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAg3yJ8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssIqKj1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OqQL3MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGZpJU9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XhxdPGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLoqfDBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tE2OtmA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IeOllJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnGwMFWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3yVTYqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVJ6JtCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocJ4NL7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deF601dEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3Sb60YDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vfQ0cXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4fVYLZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0DBT0k1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5JNCdOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtUF13ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3taKoYQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1R6M9dcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzrbCbyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZeWuDUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9tIl0QYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzGkzuXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1IQvY3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qH2q9qPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixMkFKoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgzvpJrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiTTHjI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJhK74QhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oK8LSTjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func br7kerSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxdpyC5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJvSaIwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6pv1IipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyIaltPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ur6kBbSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elipFNulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAHRjFZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YA1P0XXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfOkwHasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ttCcvr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeJGMlcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9WLVPrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FnjA4qFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTwCddIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eoLpAroYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e03HW3YYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nrnuc55jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKkQHRrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abLrNypoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXcvFdjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfAG07OjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQr86FbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lb1O0ACWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jH89raCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMBnAFGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRtA0kOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkSWL2VTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7B6xPEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tj0vte5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3swS0mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlYtaFCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VX0U3704Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVvtsATeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNvF7XeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITjtHXtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMOpIbJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57mP47pnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5fGRGKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaWmJRU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdR3nFnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSwTFnGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxAyZkENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLmbFaydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKzV2GwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwZ7oy0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpzFx40ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RQV80cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j69ivVVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAND1t57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZrnlwanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PnudZTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEcniRwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbFfFd1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnD8Lg9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3EYsxB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waMQ1z4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oz9AgxfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkxQzjWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omccgD45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grLmYOejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIwxpNAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzZwwglvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbcuP7gBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJc7NkgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPjgDinEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Py4qfJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLcEKkNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQuC5U0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCpSFAzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiN7uXJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHB5l0K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzqC4KvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWmaeA7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5JoctH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ay5NXumVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcKx4VkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zBC5JpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrIsoJZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fSs5mAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VNb3LBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeTEyKuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWIh0FF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9porJjRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvfZSjRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xjMl1RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4j9oUtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENtNIHfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXCizDl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAIR9KaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9pbBUrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3RmKJqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYVKxRlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEKJn4VAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNSonV1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq8RK10QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZYBOrgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmsaIaUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5uVzETsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqQaEPmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NU0cRXNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB0UqDMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMVN0A1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39tElcGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdD8kyu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gT0ZI7N3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMzHibXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXposvrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4j6wONaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3kT3n8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdLsXiBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJwHlNZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1AodgddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocaDlz16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icJtdmbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNEBW6zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pexop9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNKR6SpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmKbvjZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWBgOTMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSLlVkeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dijWQlPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKwE6igMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcTr6VmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y49jKWk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEGvC7UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYNoev8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6gZIUoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rafxb7EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp18jRSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoZSYJFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAwK5KBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g74yLIcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxUCx1UJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRbbpsAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z25tb8CjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgEEMajbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRWjtB5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPjTiverWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2enZTOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaO0fRhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNV1QAg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Abw9hMToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtRLXiEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlM5D20hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmYrvsDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuHbW5dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9ma6G5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJYXGPbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxOHn8boWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCxR8eExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7f2Vlo26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2DokszGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vt2G53m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuuvLTYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uG3cnlkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5yQQPAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGu9BbLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXYuup3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lw2VuvCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3r0XikU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzddTX9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZBq5cPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGf78PmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mv9H6namWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DVRyOrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62kC5Th0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJLwEeOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNXXI3CsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func up9rMdvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0UhnVSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88JAq6iGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPnDJFkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BS0frwzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Igrnf5YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MU9BGAawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwEF6LnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ifeo996jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJYjVzfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Yuvb5f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Y3BxlJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6grHdqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bukOosOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbwsLSkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERNvgjjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjjS3u1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5WrVh7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrBshhbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4A8CW4cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MLZl95AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xv3YnOKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McVeDgWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UXgXzJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQYe5bsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z18LpEweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FQOiQjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4o0otWEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k1FfHSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2itTCFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7M4WF5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMwQzTtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QCOurWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZ6XeJQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diiVRwDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xj996jXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtQi9zJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qs7YApSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLGenqHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3hXGECQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGOhA6g5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4V5HJIjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKrqTJUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hM1Hmt8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJOg6JbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoYZPGn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzTYSH1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsZJZ5XjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ScxrNBGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Pt9C5lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P91OZVC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6J90txdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfAfkGKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBnNTStCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AyqNAH7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxsJYBk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57UTe5oYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lrudb0sbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auESjeSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmK3lOjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTFkU60rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20jMAM8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjUNA9CiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Frfc9ZyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0kUuMslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RIPUrEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7JojWBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCbpcVojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJ146G3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCdSqc2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imwllxJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nm3S8uCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbvzDL9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDqg8QG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqtIFc1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gg7PMYXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYrkZdFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tOoQrJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMIwayXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkPGHkgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5T6XOUHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL2FodP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZjw7jkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnHCqHAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m16X88sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpkkXZ2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0EuE7EnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLYmlHv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67pRYWM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNFmDG9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcsKTaszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP8ef9BlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kL8eaCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bIkRCsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgz5YgGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdFz1GuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbMk3F34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZjGOzfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnT03nmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AHg7BDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33edes4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxvU2PIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmjqh8xiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjkAo3IJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NAfvF6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8oZ4oegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MO3KFtL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oAhlCgSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rjoTwb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XE3THXkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4i4U8DYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6UmGxOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdzNnVBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ5hdwOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3s4PxOB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s259USeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2eGotyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoSA8VruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cstbN0neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuDTK8DOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGFX9kozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vg8OFuonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttYlcQaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjunMmQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QDxIehUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1Ahc7uSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVeLEPqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8I56gYnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyzEDD8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWTjBanoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i55i57QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKNDvc0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXt1QG5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYjyLM1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tghG2xplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgzUo4EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ueskYWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkdecukcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWuMq5j7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znyTdREYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGBXQm7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyCLqVVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQf8GkqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbvWHER7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vqiFGElWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uogD3zdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GQREa6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvfwZ6usWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaJrpFguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMEbhuGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU1Czb3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFgqU46QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlXkFi6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QDVLsydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtrp7WDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ji5QO2djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI3AynssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1Ppl3i2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRjaEqy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvi6qKUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJFxEsyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dunBegzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VWfKQRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Li72KOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ii4byc5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Fh8ENprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33ucXZd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT2OyhEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NfyfKc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gl4hGrbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83c86aP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func effluBc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiyBd1aAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAys6GavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQEZMMTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFKc5uT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKx7wGxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qdKkPGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0ikh4OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyCoH1A6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzD88r0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6I2L73uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3CU5sEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7GVntuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1OQDyMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoJziEySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bG5p1PkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQLl8SZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9Fs8q0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWZoAMefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdgtFgcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvvF4zKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9vmEIRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqpGGhepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDEKWILkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuaJRXqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1HACu0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MOjJSOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTdOvaeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WK0pNDj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiYcfW84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwTriUM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sko7XkiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nyz5bUJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEOlfh9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeMsL2j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJZm5MM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WyWy1P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBacnRcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PI2vhnvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiwMCaRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcS9JpslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWIUTgYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiDyFSoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaBcURWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrR3IcrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgN4tJp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCs2xwDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dlzQb8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5fDg9uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zcw7YSRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFZmygnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIbR8RMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQCxONv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpYCsqMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyQXdCYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDH5Qm5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ0r4sTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmuernimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVxWUpgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lTwb1hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTR1sggfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFcao6HUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YW3LUqm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVdKiMGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QmiWqyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AafWEI9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvoJ98RGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8ExXNq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEWQEDRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X18J9ubxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPFh2FnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzyUkFwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YldXQYAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbMqXBJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLDQN7xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aKcvrdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFl06cVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func en8CCUTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwPqifwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q32CPMxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGnZgp4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQR8hiTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXT5oievWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAkXztTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNC52BAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlueC32cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKf2EZuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlKZehXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQAlE6h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBEpjfn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyWJMOacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ853dYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGyfgrt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGznOUxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBVmGT9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JC7LAfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmTPXoV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPfft6DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GHwxiBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbTIDY63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R84gTZxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDEh7TWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIILcFNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6ArblSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqohngj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4CZA4YJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMq3J3IoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aw5xqd6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqeyfQYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3JtgYwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJcQCjThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKTfc5gyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4EqbwCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeS3XNimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rmx0Wau6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AFl16fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLLOFP1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT3MPHagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDdaosAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7XgevHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3mBfGODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0e4CZJTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJxIJV5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQgFylcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVHUcl0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdalplj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6zJkMlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awU2ksF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Z9MeHz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EC4o431Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgaN8l6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGs3WTO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21pAbEYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eeyCgDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZu9x5cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDT5sscvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0C4sd9B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NqaBXcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEXUfsWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDgmL2BmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxJIFAFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3To9hEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IamUU5iMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LAQLNxg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSYS5hEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJnS35nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ullYyTDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBnFljQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xk7sjf9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoN8wwPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncuntgjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WnZhcyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CFaBs6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbSd5UF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psoR0unAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2reHWAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3humLBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAh2ItzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF6an4k5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyMOX5OiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVBDHZ9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOKteMeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9Dhjw3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMi436kOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI8mVdxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

