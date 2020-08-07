import argparse, json, os

from collections import defaultdict

def merge(input_fns):

    loc = {}  # Helper data structure
    authors = set()  # All authors
    tss = set()  # All timestamps
    for fn in input_fns:
        print('Reading %s' % fn)
        data = json.load(open(fn))
        locr = defaultdict(defaultdict)
        for i, a in enumerate(data['labels']):
            authors.add(a)
            locr[a] = {}
            for j, t in enumerate(data['ts']):
                tss.add(t)
                locr[a][t] = data['y'][i][j]
        loc[fn] = locr

    authorss = sorted(authors)  # Authors, sorted
    tsss = sorted(tss)  # Timestamps, sorted

    merged = [[0 for j in range(len(tsss))] for i in range(len(authorss))]

    for i, r in enumerate(loc):
        # print("repo: ", r)
        for j, a in enumerate(authorss):
            # print("  ", a)
            l = 0
            for k, t in enumerate(tsss):
                # print(r, a, t)
                if a in loc[r].keys():
                    if t in loc[r][a].keys():
                        l = loc[r][a][t]
                        # print("l = ", l)
                merged[j][k] = merged[j][k] + l

    fn = os.path.join('.', 'mergedAuthors.json')
    print('Writing merged authors data to %s' % (fn))
    f = open(fn, 'w')
    json.dump(
        {
            'y': merged,
            'ts': [t for t in tsss],
            'labels': [a for a in authorss]
        }, f)
    f.close()

def merge_cmdline():
    parser = argparse.ArgumentParser(description='Proov')
    parser.add_argument('input_fns', nargs='*')
    kwargs = vars(parser.parse_args())

    merge(**kwargs)

if __name__ == '__main__':
    merge_cmdline()