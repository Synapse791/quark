# Quark

A small tool to replace strings in files using environment variables.

## Why?
Quark was created to set config values inside docker containers at runtime. This means that you can use the same container in all environments by just passing it different environment variables when you run it. This enables applications that don't take their configuration values from environment variables to be deployed in the same way as applications that do.

## Installation
Just download the latest release binary from the [releases page](https://github.com/Synapse791/quark/releases/latest)! If you need Quark to be executable globally, run the following commands:

```sh
curl -sSL https://github.com/Synapse791/quark/releases/download/1.0/quark | sudo tee /usr/local/bin/quark
sudo chmod 755 /usr/local/bin/quark
```

## Basic Usage

**Some example usages:**

```
# Basic run
$ ./quark

# Custom prefix and delimeter
$ ./quark -p "MYAPP_" -d '|'

# Show help text
$ ./quark -h
```

## Environment Variable Format
Say you want to replace the string `QUARK_DB_USER` with the string `admin` in the file `/srv/myapp/config.yml`. All of this information is stored in the environment variable as follows:

```sh
export QUARK_DB_USER=/srv/myapp/config.yml:admin
             ^                 ^          ^  ^
search string|                 |          |  |
                      file path|          |  |
                                 delimeter|  |
                           replacement string|
```

Here the environment variable name is the search string, the first part of the value is the file path and the second part of the value, separated by the delimeter, is the replacement string.

## Example

This is our initial file. We could build this into a docker container.

**Initial file**
```yaml
# /srv/myapp/config.yml
myapp:
  database:
    host: QUARK_DB_HOST
    user: QUARK_DB_USER
    password: QUARK_DB_PASSWORD
    port: QUARK_DB_PORT
```

Start by exporting your environment variables. If using with docker, use the `-e` flag when running the container:

```sh
export QUARK_DB_HOST='/srv/myapp/config.yml:db.myapp.com'
export QUARK_DB_USER='/srv/myapp/config.yml:admin'
export QUARK_DB_PASSWORD='/srv/myapp/config.yml:letmein'
export QUARK_DB_PORT='/srv/myapp/config.yml:3306'
```

Now just run Quark! By default, Quark will look for environment variables starting with `QUARK_` but this can be customised by using the `-p [PREFIX]` flag. The default delimeter is set to `:` but again, this can be altered using the `-d '[DELIMETER]'` flag.

```
root@1173111040dd:/# /mnt/src/github.com/Synapse791/quark/quark
+  starting quark...
+  searching for environment variables starting with 'QUARK_'...
+  found 4 environment variables!
+  processing 4 unique keys in '/srv/myapp/config.yml'...
+  replacing 'QUARK_DB_HOST' in '/srv/myapp/config.yml' with 'db.myapp.com'...
+  replacing 'QUARK_DB_USER' in '/srv/myapp/config.yml' with 'admin'...
+  replacing 'QUARK_DB_PASSWORD' in '/srv/myapp/config.yml' with 'letmein'...
+  replacing 'QUARK_DB_PORT' in '/srv/myapp/config.yml' with '3306'...
+  successfully processed 4 environment variable in 1 files!
```

Quark will run and replace the search strings with the replacement strings to produce the following file.

**File after running Quark**
```yaml
# /srv/myapp/config.yml
myapp:
  database:
    host: db.myapp.com
    user: admin
    password: letmein
    port: 3306
```

## License
Quark is open-sourced software licensed under the [MIT License](https://opensource.org/licenses/MIT).
