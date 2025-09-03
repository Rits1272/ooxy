# Ooxy — A Bridge for Raw IP over HTTP/S

**Inspired by Cloudflare’s Oxy**, **Ooxy** is an early-stage, hackable tool that enables forwarding raw IP traffic (Layer 3 protocols) over HTTP/S proxies like mitmproxy.

Ooxy fills the "transparent proxy" gap - ideal when devices can’t speak the same network protocol but need to communicate through restrictive networks.

> * Note: Still a work-in-progress (WIP), but functional enough for experimentation.*

---

##  Why Ooxy?

- **Real-world problem**: Many networks (corporate, industrial, hotspots) only allow HTTP/S traffic, blocking everything else (SSH, MQTT, RTSP, custom TCP/UDP tools).
- **Simple bridge**: Ooxy wraps that raw IP traffic inside allowed HTTP/S channels. No need for VPNs or protocol-specific hacks.
- **Free & hackable**: Unlike closed-source offerings, Ooxy is open source and designed for tinkering and extension.

---

##  Use Cases

| Scenario | Benefit |
|----------|---------|
| SSH through a proxy | Developers can SSH into servers from within locked-down networks. |
| IoT gateways | Devices using MQTT, RTSP, or custom protocols can tunnel through HTTP/S. |
| Network debugging | Forward tools like `ping` or custom TCP services over HTTP/S for testing. |

 
