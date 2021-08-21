import slugify from "slugify";

export default function generateSlug(str: string): string {
  return slugify(str, {
    lower: true
  });
}