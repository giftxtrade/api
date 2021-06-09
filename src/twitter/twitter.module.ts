import { Module } from '@nestjs/common';
import { TwitterController } from './twitter.controller';
import { TwitterService } from './twitter.service';
import { TwitterStrategy } from './twitter.strategy';

@Module({
  controllers: [TwitterController],
  providers: [TwitterService, TwitterStrategy]
})
export class TwitterModule { }
