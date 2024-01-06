export interface HeaderItemInterface {
  href: string;
  text: string;
  title: string;
}

export const HeaderItems: HeaderItemInterface[] = [
  { text: "Home", title: "Home page", href: "/" },
  { text: "Contact", title: "Contact page", href: "/contact" },
];
