"use client";

import Link from "next/link";
import { HeaderItemInterface } from "@/content/header";
import { cn } from "@/lib/utils";
import { useSelectedLayoutSegment } from "next/navigation";

export const HeaderItem = ({ href, text, title }: HeaderItemInterface) => {
  const segment = useSelectedLayoutSegment();
  return (
    <Link title={title} href={href} legacyBehavior passHref>
      <span
        className={cn(
          "flex select-none cursor-pointer items-center text-lg font-medium transition-colors hover:text-foreground/80 sm:text-sm text-foreground/60",
          href.startsWith(`/${segment}`)
            ? "text-foreground"
            : "text-foreground/60",
        )}
      >
        {text}
      </span>
    </Link>
  );
};
