## Ooxy

Inspired by Cloudflare's [Oxy](https://blog.cloudflare.com/introducing-oxy/) framework, Ooxy is a go package that provides network connectors for the internet. Ooxy can be used as a bridege to handle communicatation between different network layers stack.

Ooxy can be used to proxy traffic with an array of protocols and can be used for various usecases - transparent proxy (with support of CONNECT request) between devices talking in different network protocols, traffic analysis, firewall

Atm, actively working on layer 3 to layer 7 connector. This could be used to proxy traffic for layer 3 device to layer 7 proxy like MITM etc. Also includes support of CONNECT request so the upstream proxies know to forward request to the destination server without trying to intercept on their own
