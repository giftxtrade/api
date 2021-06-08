import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { GoogleModule } from './google/google.module';

@Module({
  imports: [
    TypeOrmModule.forRoot(),
    UsersModule,
    GoogleModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
