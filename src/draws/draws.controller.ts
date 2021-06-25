import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { DrawsService } from './draws.service';
import { CreateDrawDto } from './dto/create-draw.dto';
import { UpdateDrawDto } from './dto/update-draw.dto';

@Controller('draws')
export class DrawsController {
  constructor(private readonly drawsService: DrawsService) {}

  @Post()
  create(@Body() createDrawDto: CreateDrawDto) {
    return this.drawsService.create(createDrawDto);
  }

  @Get()
  findAll() {
    return this.drawsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.drawsService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateDrawDto: UpdateDrawDto) {
    return this.drawsService.update(+id, updateDrawDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.drawsService.remove(+id);
  }
}
