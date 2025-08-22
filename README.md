## Ooxy  

Ooxy is a Go package inspired by Cloudflare’s [Oxy](https://blog.cloudflare.com/introducing-oxy/) that provides **network connectors across protocol layers**. It’s designed as a bridge to enable communication between different parts of the network stack.  

With Ooxy, you can:  
- Proxy traffic across multiple protocols  
- Build transparent proxies (with `CONNECT` support)  
- Analyze traffic or implement custom firewalls  
- Bridge devices that “speak” different layers  

Currently, Ooxy focuses on a **Layer 3 → Layer 7 connector**, useful for scenarios like:  
- Proxying traffic from a Layer 3 device through a Layer 7 proxy (e.g., MITM)  
- Supporting `CONNECT` requests so upstream proxies forward traffic correctly without intercepting  
