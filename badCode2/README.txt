We received the following report from a build server, where the unit tests fail:

============== snip ============== 

*** running go unit tests ***
TestRegisterWriter/Register_on_filter_with_binary_overlap
PASS: TestRegisterWriter/Register_on_filter_with_binary_overlap (0.00s)

panic: runtime error: invalid memory address or nil pointer dereference 
[signal SIGSEV: segmentation violation code=0x2 addr=0x20 pc=0x1023a07e8]

goroutine 42 [running]:
badCode2/monitor. (*monitor).newFilter.func1()
	/Users/rha/go/src/onboarding/concurrency/badCode2/monitor/monitor.go:106 +0x48
created by badCode2/monitor.(*monitor).newFilter
	/Users/cha/go/src/onboarding/concurrency/badCode2/monitor/monitor.go:101 +0xc0

============== snap ============== 

You might not be able to reproduce the SEGV, as the exact timing differs between platforms and system performance.
But let's not give up that easily, shall we? :-)
Run the unit tests with the -race option and (try) to find the cause of the reported SEGV.

If you find more than one race condition, look for the specific one causing the segmentation violation.

Even if you can't point out the root cause, this code offers many lessons. Take notes what bad coding catches your eye.
