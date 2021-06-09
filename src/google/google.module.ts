import { Module } from '@nestjs/common';
import { AuthModule } from 'src/auth/auth.module';
import { GoogleService } from 'src/google/google.service';
import { GoogleController } from './google.controller';
import { GoogleStrategy } from './google.strategy';

@Module({
  imports: [AuthModule],
  controllers: [GoogleController],
  providers: [GoogleService, GoogleStrategy]
})
export class GoogleModule { }
