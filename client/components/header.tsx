"use client";

import { HeaderItems } from "@/content/header";
import { HeaderItem } from "./header-item";
import Image from "next/image";
import { ModeToggle } from "./mode-toggle";

export const Header = () => (
  <header className="container z-40 bg-background">
    <div className="flex h-20 items-center justify-between py-6">
      <div className="flex gap-6 md-gap-10">
        <div className="hidden items-center space-x-2 md:flex">
          <Image alt="financify-logo" src="/logo.svg" width={24} height={24} />
          <span className="hidden font-bold sm:inline-block">Financify</span>
        </div>
        <nav className="hidden gap-6 md:flex">
          {HeaderItems.map((item, i) => (
            <HeaderItem key={i} {...item} />
          ))}
        </nav>
      </div>
      <ModeToggle />
    </div>
  </header>
);
