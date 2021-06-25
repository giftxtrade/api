import { Injectable } from '@nestjs/common';
import { CreateDrawDto } from './dto/create-draw.dto';
import { UpdateDrawDto } from './dto/update-draw.dto';

@Injectable()
export class DrawsService {
  create(createDrawDto: CreateDrawDto) {
    return 'This action adds a new draw';
  }

  findAll() {
    return `This action returns all draws`;
  }

  findOne(id: number) {
    return `This action returns a #${id} draw`;
  }

  update(id: number, updateDrawDto: UpdateDrawDto) {
    return `This action updates a #${id} draw`;
  }

  remove(id: number) {
    return `This action removes a #${id} draw`;
  }
}
