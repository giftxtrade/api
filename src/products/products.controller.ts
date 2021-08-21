import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards, Query, Request, HttpStatus, HttpException } from '@nestjs/common';
import { ProductsService } from './products.service';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { BAD_REQUEST, NOT_FOUND } from 'src/util/exceptions';
import { UNAUTHORIZED } from '../util/exceptions';

@Controller('products')
export class ProductsController {
  constructor(private readonly productsService: ProductsService) {}

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Request() req, @Body() createProductDto: CreateProductDto) {
    const { user } = req.user;
    if (user.email !== 'moahammedayaan.dev@gmail.com')
      throw UNAUTHORIZED('You are not authorized to perform this action');

    return await this.productsService.create(createProductDto);
  }

  @Get()
  async findAll(
    @Query('limit') limit: number = 50,
    @Query('page') page: number = 1,
    @Query('min_price') minPrice: number,
    @Query('max_price') maxPrice: number,
    @Query('search') search: string,
    @Query('sort') sort: string,
  ) {
    const prevPage = page - 1;
    const results = await this.productsService
      .findAllWithLimit(
        limit,
        prevPage * limit,
        minPrice,
        maxPrice,
        search ? search.trim() : undefined,
        sort ? sort.trim().toLowerCase() : undefined
      );

    // If result is empty then check assume search is a product key or URL
    // if no products are found then throw HTTP Exception
    if (results.length === 0) {
      const productKey = this.amazonProductKeyFromSearchOrUrl(search.trim())
      const productFromKey = await this.productsService.findByProductKey(productKey);
      if (productFromKey && page == 1) {
        return [productFromKey];
      } else if (!productFromKey && page == 1) {
        const url = `https://amazon.com/s?k=${productKey}`;

        // TODO: Find product by Amazon product key search

      }

      throw BAD_REQUEST('No results');
    }
    return results;
  }

  @Get(':id')
  async findOne(@Param('id') id: string) {
    const product = await this.productsService.findOne(+id);
    if (!product)
      throw NOT_FOUND('Product not found')
    return product;
  }

  private amazonProductKeyFromSearchOrUrl(search: string) {
    // Check if search is a URL
    if (search.includes('amazon.com') && search.includes('dp')) {
      // Assume productKey is a URL and then split by '/'
      const splitBySlash = search.split('/');
      // Amazon product URL format: https://www.amazon.com/dp/[PRODUCT_KEY]...
      // So if 'dp' is found then this is a valid products URL
      const dpIndex = splitBySlash.findIndex(v => v === 'dp');

      if (dpIndex > -1) {
        const key = splitBySlash[dpIndex + 1]; // [..., 'dp', 'KEY?...']
        return key.split('?')[0];
      }
    }
    return search;
  }
}
