<script>
// COMPONENTS
import BrokenImage from '@/assets/interface-icons/broken-icon.svg';

// UTILS
import ErrorHandler from '@/utils/ErrorHandler';
import EmojiUnicodeRegex from '@/utils/EmojiUnicodeRegex';
import emojiLookup from '@/utils/emojis.json';
import { asciiHash } from '@/utils/MiscHelpers';

// CONSTANTS
import { faviconApi as defaultFaviconApi, faviconApiEndpoints, iconCdns } from '@/utils/defaults';

// ICON LIBRARY (prefer ESM import for optimal tree-shaking and static analysis)
import simpleicons from 'simple-icons';

export default {
  name: 'Icon',
  props: {
    icon: { type: String, required: true },
    url: { type: String, default: '' },
    size: { type: String, default: 'medium', validator: s => ['small', 'medium', 'large'].includes(s) }
  },
  components: { BrokenImage },
  // ...rest of your component logic
};
</script>
