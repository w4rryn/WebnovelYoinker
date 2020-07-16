# WebnovelYoinker
Downloads and converts webnovels to Epub or PDF

#### Currently supported websites:
  - wuxiaworld.com (wuxia)
  - crimsonmagic.me (crimsonmagic)

#### Currently supported export formats:
  - Epub (epub)
  
#### Basic usage:
Declaring volumes as YAML file:

```yaml
--- # example-books.yaml
- Author:   Park Saenal
  CoverImageURL: https://cdn.wuxiaworld.com/images/covers/og.jpg
  Language: English
  Title:    Overgeared Volume 1
  Year:     2013
  Website : wuxia
  ExportFormat : epub
  ChapterURLs:
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-1
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-2
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-3
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-4
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-5
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-6
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-7
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-8
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-9
    - https://www.wuxiaworld.com/novel/overgeared/og-chapter-10

- Title:    Konosuba Volume 1
  Author:   Natsume Akatsuki
  CoverImageURL:    https://i.imgur.com/eLIkho2.png
  Language: English
  Year:     2013
  Website : crimsonmagic
  ExportFormat : epub
  ChapterURLs:
    - https://www.crimsonmagic.me/archive/gifting-1-p/
    - https://www.crimsonmagic.me/archive/gifting-1-1/
    - https://www.crimsonmagic.me/archive/gifting-1-2/
    - https://www.crimsonmagic.me/archive/gifting-1-3/
    - https://www.crimsonmagic.me/archive/gifting-1-4/
    - https://www.crimsonmagic.me/archive/gifting-1-e/
```
Downloading declared books:
```zsh
Linux and Windows (powershell):
goyoinker scrape -in example-books.yaml -out OUTPUT_DIRECTORY -r 3
```
The flag -r declares how many books should be downloaded in parallel 
