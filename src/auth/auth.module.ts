import { Module } from '@nestjs/common';
import { UsersModule } from 'src/users/users.module';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { JwtModule } from '@nestjs/jwt';
import { PassportModule } from '@nestjs/passport';
import { JWT } from '../../auth-tokens.json';

@Module({
  imports: [
    UsersModule,
    PassportModule,
    JwtModule.register({
      secret: JWT.SECRET,
      signOptions: { expiresIn: '60s' },
    })
  ],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [AuthService]
})
export class AuthModule {}