package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.infrastructure.cache.redis.RedisConfig;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.jdbc.autoconfigure.DataSourceAutoConfiguration;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Import;

@TestConfiguration
@Import({RedisConfig.class})
@EnableAutoConfiguration(exclude = {DataSourceAutoConfiguration.class})
public class RedisTestConfiguration {
}
