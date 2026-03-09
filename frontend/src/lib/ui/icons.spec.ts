import { describe, expect, it } from "vitest";
import { icons } from "lucide";
import { iconMap } from "./icons";

describe("iconMap", () => {
  it("maps each app icon key to a Lucide icon that exists", () => {
    for (const lucideKey of Object.values(iconMap)) {
      expect(icons).toHaveProperty(lucideKey);
    }
  });
});
