# Test cases — Centered Hello Word page

Risk level: low. Single public landing page, no auth, one read-only value, plain rendering.

## Cases

**Scenario**: Show stored greeting on centered page
**Given**: PostgreSQL greeting row exists with text `Hello Word`, and backend `/v1/greeting` returns that stored value
**When**: Guest opens landing page
**Then**: Page displays exactly `Hello Word`, centered horizontally and vertically on single screen
**Check**: render_url

**Scenario**: Greeting is not hardcoded in frontend page copy
**Given**: Backend returns stored greeting value, and frontend page source is inspected after load
**When**: Guest inspects rendered page source or network response
**Then**: Frontend page copy does not contain hardcoded greeting text; displayed text comes from backend response
**Check**: fetch_url

**Scenario**: Plain white background, black text, no extra chrome or animation
**Given**: Approved design loaded and page rendered with stored greeting
**When**: Guest views landing page
**Then**: Computed background is `#FFFFFF`, computed text color is `#000000`, and no extra chrome or animation is present
**Check**: measure_styles

**Scenario**: Empty stored greeting does not render empty message
**Given**: Greeting row exists with empty text
**When**: Guest opens landing page
**Then**: Empty greeting is not rendered as visible page content
**Check**: render_url

**Scenario**: Short plain greeting renders unchanged within centered text line
**Given**: Greeting row exists with short plain string `Hi`
**When**: Guest opens landing page
**Then**: Page displays `Hi` unchanged in same centered text region
**Check**: render_url

**Scenario**: Missing greeting row shows no loading, empty, or error state in approved design
**Given**: Greeting row is missing
**When**: Guest opens landing page
**Then**: Page does not show loading copy, empty state copy, or error state copy from approved design
**Check**: render_url

**Scenario**: Guest has no permission gate
**Given**: Guest is unauthenticated
**When**: Guest opens landing page
**Then**: Page remains public and loads without auth prompt or redirect
**Check**: fetch_url

**Scenario**: Latest committed greeting value shows after row changes during read
**Given**: Greeting row changes from one value to another before refresh
**When**: Guest reloads page after change is committed
**Then**: Page shows latest committed stored value
**Check**: render_url

**Scenario**: Backend or database unavailable shows no loading, empty, or error state in approved design
**Given**: Backend API or PostgreSQL is unavailable
**When**: Guest opens landing page
**Then**: Page does not render loading, empty, or error state from approved design
**Check**: render_url
