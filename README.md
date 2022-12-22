# Wake On Lan

![Golang](https://img.shields.io/github/actions/workflow/status/starudream/wake-on-lan/golang.yml?label=golang&style=for-the-badge)
![Docker](https://img.shields.io/github/actions/workflow/status/starudream/wake-on-lan/docker.yml?label=docker&style=for-the-badge)
![Release](https://img.shields.io/github/v/release/starudream/wake-on-lan?include_prereleases&sort=semver&style=for-the-badge)
![License](https://img.shields.io/github/license/starudream/wake-on-lan?style=for-the-badge)

## Usage

```
Usage of ./bin/app:
  -addr string
    	broadcast ip (default "255.255.255.255")
  -mac string
    	mac address
  -port string
    	 (default "9")
```

### Docker

![Version](https://img.shields.io/docker/v/starudream/wake-on-lan?sort=semver&style=for-the-badge)
![Size](https://img.shields.io/docker/image-size/starudream/wake-on-lan?sort=semver&style=for-the-badge)
![Pull](https://img.shields.io/docker/pulls/starudream/wake-on-lan?style=for-the-badge)

```bash
docker pull starudream/wake-on-lan
```

```bash
docker run --rm --network host starudream/wake-on-lan /app -mac 00:00:00:00:00:00
```

## License

[Apache License 2.0](./LICENSE)
