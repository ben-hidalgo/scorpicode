for pid in $(ps aux | grep "go-build.*main" | grep -v "grep" | awk '{print $2}'); do kill $pid; done
