// Code generated by tygo. DO NOT EDIT.

//////////
// source: dto.go

export interface Response {
  message: string;
}
export interface Result {
  data: any;
}
export interface Errors {
  errors: string[];
}
export interface DeleteStatus {
  deleted: boolean;
}
export interface User {
  id: number /* int64 */;
  name: string;
  email: string;
  imageUrl?: string;
  active: boolean;
  phone?: string;
  admin?: boolean;
}
export interface CreateUser {
  name: string;
  email: string;
  imageUrl?: string;
  phone?: string;
}
export interface Auth {
  user: User;
  token: string;
}
export interface Category {
  id: number /* int64 */;
  name: string;
  description?: string;
}
export interface CreateCategory {
  name: string;
  description?: string;
  url?: string;
}
export interface Product {
  id: number /* int64 */;
  title: string;
  description?: string;
  productKey: string;
  imageUrl: string;
  totalReviews: number /* int32 */;
  rating: number /* float32 */;
  price: string;
  currency: string;
  url: string;
  categoryId?: number /* int64 */;
  category?: Category;
  createdAt: string /* RFC3339 */;
  updatedAt: string /* RFC3339 */;
  origin: string;
}
export interface CreateProduct {
  title: string;
  description?: string;
  productKey: string;
  imageUrl?: string;
  rating: number /* float32 */;
  price: string;
  originalUrl: string;
  totalReviews: number /* uint */;
  category: string;
}
export interface ProductFilter {
  search?: string;
  limit: number /* int32 */;
  page: number /* int32 */;
  minPrice?: number /* float32 */;
  maxPrice?: number /* float32 */;
  sort?: string;
}
export interface Participant {
  id: number /* int64 */;
  name: string;
  email: string;
  address?: string;
  organizer: boolean;
  participates: boolean;
  accepted: boolean;
  eventId?: number /* int64 */;
  event?: Event;
  userId?: number /* int64 */;
  user?: User;
}
export interface CreateParticipant {
  email: string;
  name?: string;
  address?: string;
  organizer?: boolean;
  participates?: boolean;
}
export interface Event {
  id: number /* int64 */;
  name: string;
  slug?: string;
  description?: string;
  budget: string;
  invitationMessage?: string;
  drawAt: string /* RFC3339 */;
  closeAt: string /* RFC3339 */;
  createdAt: string /* RFC3339 */;
  updatedAt: string /* RFC3339 */;
  participants?: Participant[];
}
export interface CreateEvent {
  name: string;
  description?: string;
  budget: number /* float32 */;
  inviteMessage?: string;
  drawAt: string /* RFC3339 */;
  closeAt: string /* RFC3339 */;
  participants?: CreateParticipant[];
}
export interface UpdateEvent {
  name?: string;
  description?: string;
  budget?: number /* float32 */;
  drawAt?: string /* RFC3339 */;
  closeAt?: string /* RFC3339 */;
}
