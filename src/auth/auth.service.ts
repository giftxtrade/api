import { Injectable } from '@nestjs/common';
import { User } from 'src/users/entities/user.entity';
import { UsersService } from 'src/users/users.service';
import { JwtService } from '@nestjs/jwt';
import { CreateUserDto } from '../users/dto/create-user.dto';
import JwtPayload from 'src/util/jwtPayload';

@Injectable()
export class AuthService {
  constructor(
    private readonly userServices: UsersService,
    private readonly jwtService: JwtService
  ) { }

  async validateUser(email: string): Promise<User> {
    return this.userServices.findOne(email);
  }

  async login(createUser: CreateUserDto, gToken: string): Promise<{ user: User, accessToken: string }> {
    const user = await this.userServices.findOneOrCreate(createUser);
    const payload: JwtPayload = { user, gToken };
    return {
      user: user,
      accessToken: this.jwtService.sign(payload),
    };
  }
}
