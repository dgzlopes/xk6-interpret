# xk6-interpret

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a TOY extension. USE AT YOUR OWN RISK! OR... DON'T USE IT! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/dgzlopes/xk6-interpret@latest
  ```

## Example

```javascript
import interpret from 'k6/x/interpret';

export default function() {
    var result = interpret.run(
    `package interpret
    
    import "fmt"


    // fibonacci is a function that returns
    // a function that returns an int.
    func fibonacci() func() int {
        x, y := 0, 1
        return func() int {
            x, y = y, x + y
            fmt.Println("Calculated:",x)
            return x
        }
    }

    func Run(s interface{}) interface{} {
        f := fibonacci()
        var results []int
        var i int64
        for i = 0; i < s.(int64); i++ {
            results = append(results,f())
        }
        return results
    }
    `,
    10
    );

    console.log(`Results: ${result}`)
}
```

Result output:

```
$ ./k6 run script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

Calculated: 1
Calculated: 2
Calculated: 4
Calculated: 8
Calculated: 16
Calculated: 32
Calculated: 64
Calculated: 128
Calculated: 256
Calculated: 512
INFO[0000] Results: 1,2,4,8,16,32,64,128,256,512         source=console

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

    data_received........: 0 B 0 B/s
    data_sent............: 0 B 0 B/s
    iteration_duration...: avg=1.16ms min=1.16ms med=1.16ms max=1.16ms p(90)=1.16ms p(95)=1.16ms
    iterations...........: 1   66.559293/s
```
