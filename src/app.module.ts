import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { GoogleModule } from './google/google.module';
import { TwitterModule } from './twitter/twitter.module';

@Module({
  imports: [
    TypeOrmModule.forRoot(),
    UsersModule,
    GoogleModule,
    TwitterModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
