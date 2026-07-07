package com.example.alertbridge.alerts.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.JacksonJsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;
import tools.jackson.databind.ObjectMapper;

@Configuration
public class RedisConfig {
    @Bean
    RedisTemplate<String, AlertCurrentStateRedis> alertRedisTemplate(RedisConnectionFactory connectionFactory,
                                                                     ObjectMapper objectMapper) {

        RedisTemplate<String, AlertCurrentStateRedis> template = new RedisTemplate<>();

        template.setConnectionFactory(connectionFactory);
        template.setKeySerializer(new StringRedisSerializer());

        JacksonJsonRedisSerializer<AlertCurrentStateRedis> serializer = new JacksonJsonRedisSerializer<>(
                objectMapper,
                AlertCurrentStateRedis.class
        );

        template.setValueSerializer(serializer);
        template.afterPropertiesSet();

        return template;
    }

    @Bean
    AlertCurrentStateWriterPort redisAlertCurrentStateWriter(RedisTemplate<String, AlertCurrentStateRedis> redisTemplate) {
        return new RedisAlertCurrentStateWriter(redisTemplate);
    }

    @Bean
    AlertCurrentStateReaderPort redisAlertCurrentStateReader(RedisTemplate<String, AlertCurrentStateRedis> redisTemplate) {
        return new RedisAlertCurrentStateReader(redisTemplate);
    }
}
