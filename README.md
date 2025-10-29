<div align="center">
  <a href="https://bosquesdeagua.ar/">
    <img src="images/cropped-logo-blanco-big.webp" alt="logo" width="500"/>
  </a>
  </br>
  </br>
</div>

# App Control

Bosques de Agua App Control is a web application for controlling and monitoring appliances.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/)
- [Podman](https://podman.io/)

### Installation

1.  **Clone the repository:**

    ```sh
    git clone git@github.com:masch/bda-app-control.git
    ```

2.  **Set up the database:**

    This project uses a PostgreSQL database. You can use the following `podman` commands to start a PostgreSQL container:

    - **Pull and init database container:**

      ```sh
      make db-init
      ```

    - **Test the database connection:**

      ```sh
      podman exec -it bda-db psql -U tabaquillo -d bosque
      ```

    - **Reset the database (optional):**

      ```sh
      podman stop bda-db && \
      podman rm bda-db && \
      podman volume rm bda-pgdata
      ```

3.  **Set up the environment variables:**

    Create a `.env` file in the root of the project with the following content using `.env.template` file as a template:

    ```
    DB_HOST=
    DB_NAME=
    DB_USERNAME=
    DB_PASSWORD=
    VERSION=
    ```

4.  **Open Ports in Firewall (for Hotspot/LAN access):**

    On some Linux distributions like Fedora, you may need to configure the firewall to allow access to the application from other devices on the network. If you are running the application on a machine that is also acting as a Wi-Fi hotspot, you'll need to open port 4000 in the appropriate firewall zone.

    The following command will add a permanent rule to the `nm-shared` zone to allow TCP traffic on port 4000 and then reload the firewall to apply the change:

    ```sh
    sudo firewall-cmd --permanent --zone=nm-shared --add-port=4000/tcp && \
    sudo firewall-cmd --reload
    ```

    - `--permanent`: Ensures the firewall rule persists after a reboot.
    - `--zone=nm-shared`: Specifies that the rule applies to the `nm-shared` zone, which is often used for network connections shared via NetworkManager (like a hotspot).
    - `--add-port=4000/tcp`: Opens port 4000 for TCP traffic.
    - `firewall-cmd --reload`: Applies the new firewall rules immediately.

## Usage

To run the web server, execute the following command from the root of the project:

```sh
go run ./cmd/web
```

The application will be available at [http://localhost:4000/bosquesdeagua](http://localhost:4000/bosquesdeagua).

## Technologies Used

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [bun](https://bun.uptrace.dev/) - SQL-first Golang ORM

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
