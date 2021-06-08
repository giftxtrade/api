import { Module } from '@nestjs/common';
import { GoogleService } from 'src/google/google.service';
import { GoogleController } from './google.controller';
import { GoogleStrategy } from './google.strategy';

@Module({
  controllers: [GoogleController],
  providers: [GoogleService, GoogleStrategy]
})
export class GoogleModule { }
