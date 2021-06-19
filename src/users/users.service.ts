import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DeleteResult, Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { User } from './entities/user.entity';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User)
    private usersRepository: Repository<User>
  ) { }

  async insert(user: CreateUserDto): Promise<User> {
    const userEntity = new User();
    userEntity.email = user.email;
    userEntity.name = user.name;
    userEntity.imageUrl = user.imageUrl;

    return await userEntity.save();
  }

  async findAll(): Promise<User[]> {
    return await this.usersRepository.find();
  }

  async findById(id: number): Promise<User> {
    return await this.usersRepository
      .findOne({
        where: { id }
      });
  }

  async findByEmail(email: string): Promise<User> {
    return await this.usersRepository
      .findOne({
        where: { email: email }
      });
  }

  async findOne(email: string): Promise<User> {
    return await this.findByEmail(email);
  }

  async findOneOrCreate(user: CreateUserDto): Promise<User> {
    const existingUser = await this.findOne(user.email);
    if (existingUser) {
      let changed = false;
      if (user.imageUrl !== existingUser.imageUrl) {
        existingUser.imageUrl = user.imageUrl;
        changed ||= true;
      }

      if (user.name !== existingUser.name) {
        existingUser.name = user.name;
        changed ||= true;
      }

      if (changed)
        return await existingUser.save();
      return existingUser;
    }
    return await this.insert(user);
  }

  async remove(email: string): Promise<DeleteResult> {
    return await this.usersRepository.delete(email);
  }
}
