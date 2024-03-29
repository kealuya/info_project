package com.szhtjykj.speech;

import com.zaxxer.hikari.HikariDataSource;
import org.beetl.sql.core.SQLManager;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.annotation.Bean;
import org.noear.solon.annotation.Configuration;
import org.noear.solon.annotation.Inject;

import javax.sql.DataSource;

//配置数据源
@Configuration
public class Config {
    @Bean(name = "db1", typed = true)
    public DataSource db1(@Inject("${demo.db1}") HikariDataSource dataSource) {
        return dataSource;
    }

    @Bean
    public void db1m(@Db("db1") SQLManager sqlManager) {
        //sqlManager.setNc(new DefaultNameConversion());
    }
}