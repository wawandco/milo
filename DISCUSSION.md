# Rules in discussion.

This doc holds some of the rules we have in discussion before getting started with the implementation of a reviewer for these.

- [REVIEW] id-class-ad-disabled: The id and class attributes cannot use the ad keyword, it will be blocked by adblock software.
- [REVIEW] id-class-value: The id and class attribute values must meet the specified rules.
- [REVIEW] The `<script>` tag cannot be used in a `<head>` tag.
- [REVIEW] empty-tag-not-self-closed: The empty tag should not be closed by self.
- [REVIEW] tag-self-close: Empty tags must be self closed.
- [REVIEW] tags-check: Allowing specify rules for any tag and validate that
- [REVIEW] space-tab-mixed-disabled: Do not mix tabs and spaces for indentation.
- [REVIEW] spec-char-escape: Special characters must be escaped.
- [REVIEW] href-abs-or-rel: An href attribute must be either absolute or relative.