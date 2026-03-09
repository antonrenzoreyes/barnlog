import type { icons } from "lucide";

export const iconMap = {
  addAnimal: "PawPrint",
  addLogEvent: "ClipboardPlus",
  app: "Warehouse",
  back: "ChevronLeft",
  editAnimal: "Pencil",
  emptyState: "CircleHelp",
  eventFeed: "Utensils",
  eventMedication: "Pill",
  eventNote: "NotebookPen",
  eventWeight: "Scale",
  search: "Search",
  timelineTime: "Clock3",
  uploadPhoto: "Upload",
} as const satisfies Record<string, keyof typeof icons>;

export type AppIconKey = keyof typeof iconMap;
