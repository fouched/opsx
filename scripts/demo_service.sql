insert into remote_services (name, type, host, port, status, health_url,
                             command_start, command_stop, command_restart, last_checked)
values ('Remote Service 01', 'Application', 'localhost', 8080, 'Down', 'http://localhost:8080/health',
        '', '', '', CURRENT_TIMESTAMP);