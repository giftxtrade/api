import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, HttpException, HttpStatus, Query } from '@nestjs/common';
import { ProductsService } from './products.service';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';

@Controller('products')
export class ProductsController {
  constructor(private readonly productsService: ProductsService) {}

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Body() createProductDto: CreateProductDto) {
    return await this.productsService.create(createProductDto);
  }

  @Get()
  async findAll(
    @Query('limit') limit: number = 50,
    @Query('page') page: number = 1,
    @Query('min_price') minPrice: number,
    @Query('max_price') maxPrice: number,
    @Query('search') search: string
  ) {
    try {
      const prevPage = page - 1;
      return await this.productsService
        .findAllWithLimit(
          limit,
          prevPage * limit,
          minPrice,
          maxPrice,
          search
        );
    } catch (e) {
      throw new HttpException({
        message: 'Page not available'
      }, HttpStatus.BAD_REQUEST)
    }
  }

  @Get(':id')
  async findOne(@Param('id') id: string) {
    const product = await this.productsService.findOne(+id);
    if (!product) {
      throw new HttpException({
        messsage: 'Product not found'
      }, HttpStatus.NOT_FOUND)
    }
    return product;
  }
}
